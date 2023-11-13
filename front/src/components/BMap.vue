<template>
    <BMap 
        :width="width" 
        :height="height" 
        :ak="mapAk" 
        :center="location.point"
        :style="mapStyle"
        enableDragging
        enableDoubleClickZoom
        enableScrollWheelZoom
        ref="map"
        @initd="handleInitd" 
        :mapType="mapType"
        @click="handleMapClick"
    >
        <BControl v-if="enableModify" anchor="BMAP_ANCHOR_TOP_LEFT" :offset="{ x: 20, y: 20 }">
            <NButton v-if="!isFullScreen" type="primary" @click="isFullScreen = true">修改</NButton>
            <NButton v-if="isFullScreen" type="primary" @click="isFullScreen = false">完成</NButton>
        </BControl>
        <BZoom anchor="BMAP_ANCHOR_TOP_LEFT" :offset="{ x: 22, y: 60 }" />
        <BScale anchor="BMAP_ANCHOR_BOTTOM_RIGHT" />
        <BNavigation3d anchor="BMAP_ANCHOR_TOP_LEFT" :offset="{ x: 10, y: 140 }" />
        <template v-if="mode == 'single' && !disableMarker">
            <BMarker :position="{ lat: singlePoint.lat, lng: singlePoint.lng }" icon="simple_red" />
        </template>
        <template v-if="mode == 'multi'">
            <template v-if="enableMarkerRender">
                <BInfoWindow
                    show
                    v-for="(point, index) in multiPointsRender" :key="index"
                    :position="{ lat: point.lat, lng: point.lng }"
                    :title="point.title"
                ></BInfoWindow>
            </template>
            <template v-else>
                <BMarker v-for="(point, index) in multiPointsRender" :key="index"
                    :position="{ lat: point.lat, lng: point.lng }"
                    :visible="true"
                    :icon="point.ico ? {
                        imageUrl: point.ico,
                        imageSize: {
                            width: 100,
                            height: 100,
                        },
                    } : 'simple_red'"
                    @click="point.click()"
                />
            </template>
        </template>
        <BControl anchor="BMAP_ANCHOR_TOP_RIGHT" :offset="{ x: 20, y: 20 }" class="custom-control">
            <select style="border-radius: 8px;" class="select" name="" id="" v-model="mapType">
                <option value="BMAP_NORMAL_MAP">普通地图</option>
                <option value="BMAP_SATELLITE_MAP">卫星地图</option>
                <option value="BMAP_EARTH_MAP">地球模式</option>
            </select>
        </BControl>
        <BLocation anchor="BMAP_ANCHOR_BOTTOM_RIGHT"></BLocation>
        <BControl anchor="BMAP_ANCHOR_BOTTOM_LEFT" v-if="showSearchComponent" class="custom-control mx5" 
            :offset="{
                x: 0,
                y: 80
            }"
            style="width: calc(100% - 40px);"
        >
            <NInput  placeholder="搜索地名" :bordered="false" v-model:value="searchForm.keyword" @click="showSearch = true">
                <template #prefix>
                    <NIcon :component="Search"></NIcon>
                </template>
            </NInput>
        </BControl>
    </BMap>
    <NDrawer title="搜索" placement="right" v-model:show="showSearch" :width="500">
        <NDrawerContent>
            <NInputGroup>
                <NInput v-model:value="searchForm.keyword" placeholder="搜索地名" @keypress.enter="search">
                    <template #prefix>
                        <NIcon :component="Search"></NIcon>
                    </template>
                </NInput>
                <NButton type="primary" @click="search">搜索</NButton>
            </NInputGroup>
            <NSpin class="center w100 my5" v-if="isSearching" :size="90">
                <template #description>
                    搜索中
                </template>
            </NSpin>
            <NList class="mx3">
                <NListItem class="clickable" v-for="result in searchResult">
                    <NThing @click="handleSearchLocate(result.point)">
                        <template #header>
                            {{ result.title }}
                        </template>
                        <template #description>
                            {{ result.address }}
                        </template>
                    </NThing>
                </NListItem>
            </NList>
        </NDrawerContent>
    </NDrawer>
</template>
  
<script lang="ts" setup>
import { Search } from '@vicons/ionicons5'
import { NButton, NDrawer, NDrawerContent, NIcon, NInput, NInputGroup, NList, NListItem, NSpin, NThing, useMessage } from 'naive-ui'
import { onMounted, ref, PropType, watch, computed, VNodeChild } from 'vue'
import {
    BMap,
    BZoom,
    BScale,
    BNavigation3d,
    BControl,
    BMarker,
    useBrowserLocation,
    MapType,
    BLocation,
    BInfoWindow,
usePointGeocoder,
PointGeocoderResult,
} from 'vue3-baidu-map-gl'

const message = useMessage()

// @ts-ignore
const mapAk = ref((import.meta as any).env.VITE_APP_BMAP_AK)
const map = ref()
const BMapGL = ref()

const handleInitd = (e: any) => {
    BMapGL.value = e.BMapGL
    if (settingProps.center.lng == 0 || settingProps.center.lat == 0) {
        get()
    } else {
        handleLocate({
            lng: settingProps.center.lng,
            lat: settingProps.center.lat,
        })
    }
}

const { get: geocoderGet, result: geoResult, isLoading: geoLoading } = usePointGeocoder<PointGeocoderResult>()

const { get, location } = useBrowserLocation({
    enableHighAccuracy: true,
    enableSDKLocation: true,
}, () => {
    if (settingProps.initCenter) {
        map.value.resetCenter()
    }
})
const emitEvents = defineEmits<{
    (event: 'onLocate', position: { lng: number, lat: number }): void
    (event: 'onSearch', position: { lng: number, lat: number }): void
    (event: 'onClick', position: { lng: number, lat: number, name: string }): void
}>()
const handleLocate = async (position: { lng: number, lat: number }) => {
    const mapInstance = map.value.getMapInstance()
    mapInstance.setCenter(position)
    emitEvents('onLocate', position)
}
const handleSearchLocate = (position: { lng: number, lat: number }) => {
    const mapInstance = map.value.getMapInstance()
    mapInstance.setCenter(position)
    emitEvents('onLocate', position)
}
watch(location, (v) => {
    if (v.point) {
        handleLocate(v.point)
    }
}, { deep: true })

const mapType = ref('BMAP_NORMAL_MAP' as MapType)

const height = ref(550)
const width = ref(550)

const disableMarker = ref(false)
const isFullScreen = ref(false)
const mapStyle = computed(() => {
    return {
        position: isFullScreen.value ? 'fixed' : 'relative',
        top: isFullScreen.value ? '0' : 'auto',
        left: isFullScreen.value ? '0' : 'auto',
        zIndex: isFullScreen.value ? '999' : 'auto',
    }
})

const showSearch = ref(false)
const searchForm = ref({
    keyword: '',
})
const isSearching = ref(false)
const searchResult = ref<{
    title: string,
    point: { lng: number, lat: number },
    address: string,
}[]>()
const showSearchComponent = computed(() => {
    if (!settingProps.enableSearch) {
        return false
    }
    if (settingProps.searchOnEdit) {
        return isFullScreen.value
    } else {
        return true
    }
})
const search = async () => {
    if (searchForm.value.keyword.length == 0) {
        message.error('请输入搜索关键字')
        return
    }

    const mapInstance = map.value.getMapInstance()

    const local = new BMapGL.value.LocalSearch(mapInstance, {
        onSearchComplete: (e: any) => {
            isSearching.value = false
            if (e._pois && e._pois.length > 0) {
                searchResult.value = e._pois.map((item: any) => {
                    return {
                        title: item.title,
                        point: item.point,
                        address: item.address,
                    }
                })
            } else {
                message.error('未搜索到结果')
            }
        }
    })
    isSearching.value = true
    local.search(searchForm.value.keyword)
}

const isMounted = ref(false)
const waitQueue = ref([] as (() => void)[])
onMounted(() => {
    isMounted.value = true
    waitQueue.value.forEach((fn) => {
        fn()
    })
})

watch(isFullScreen, (v) => {
    const handler = () => {
        if (v) {
            message.info('请点击地图选择位置')
            height.value = window.innerHeight
            width.value = window.innerWidth
        } else {
            // get parent element through map instance
            const parent = map.value.$el.parentElement
            height.value = parent.clientHeight
            width.value = parent.clientWidth
        }
    }

    if (!isMounted.value) {
        waitQueue.value.push(handler)
    } else {
        handler()
    }
}, { immediate: true })


type MapLocationMode = 'single' | 'multi' | 'none'

const settingProps = defineProps({
    mode: {
        type: String as PropType<MapLocationMode>,
        default: 'single',
    },
    initCenter: {
        type: Boolean,
        default: true,
    },
    enableModify: {
        type: Boolean,
        default: true,
    },
    enableMarkerRender: {
        type: Boolean,
        default: false,
    },
    enableAddCenterMarker: {
        type: Boolean,
        default: true,
    },
    enableSearch: {
        type: Boolean,
        default: false,
    },
    enableGeoDecoder: {
        type: Boolean,
        default: true,
    },
    searchOnEdit: {
        type: Boolean,
        default: false,
    },
    center: {
        type: Object as PropType<{ lng: number, lat: number }>,
        default: () => {
            return {
                lng: 0,
                lat: 0,
            }
        },
    },
})

type MapCoordinatePoint = {
    lng: number
    lat: number,
    title?: string,
    detail?: any,
    ico?: string,
    renderIco?: string,
    click?: () => void,
}

const singlePoint = defineModel<MapCoordinatePoint>('point', {
    required: false,
    default: {
        lng: 0,
        lat: 0,
    },
    local: true
})

const multiPoints = defineModel<MapCoordinatePoint[]>('points', {
    required: false,
    default: [],
    local: true,
})

watch(location, () => {
    const point = location.value.point
    const lng = point.lng
    const lat = point.lat
    if (settingProps.enableAddCenterMarker) {
        if (settingProps.mode == 'single') {
            singlePoint.value.lat = lat
            singlePoint.value.lng = lng
        } else if (settingProps.mode == 'multi') {
            if (multiPoints.value.length == 0) {
                multiPoints.value.push({
                    lat: lat,
                    lng: lng,
                })
            }
        }
    }
}, { deep: true })

const isClickMarker = ref<boolean>(false)
const multiPointsRender = computed(() => {
    return multiPoints.value.map((point) => {
        return {
            ...point,
            click () {
                isClickMarker.value = true
                setTimeout(() => {
                    isClickMarker.value = false
                }, 100)

                point.click && point.click()
            }
        }
    })
})

const handleMapClick = async ({
    target, latlng
}: {
    target: HTMLDivElement,
    latlng: { lat: number, lng: number },
}) => {
    if (isClickMarker.value) {
        return
    }


    let name = ''
    if (settingProps.enableGeoDecoder) {
        geocoderGet(latlng)
        for (let i = 0; i < 30; i++) {
            if (!geoLoading.value) {
                name = (geoResult.value?.address || '') + (geoResult.value?.surroundingPois[0]?.title || '')
                break
            }
            await new Promise((resolve) => setTimeout(() => resolve(null) , 100))
        }
    }
    
    emitEvents('onClick', {
        lng: latlng.lng,
        lat: latlng.lat,
        name: name,
    })

    if(!isFullScreen.value) {
        return
    }

    if (settingProps.mode == 'single') {
        singlePoint.value.lat = latlng.lat
        singlePoint.value.lng = latlng.lng
    } else if (settingProps.mode == 'multi') {
        multiPoints.value.push({
            lat: latlng.lat,
            lng: latlng.lng,
        })
    }

    disableMarker.value = true
    setTimeout(() => {
        disableMarker.value = false
    }, 0)
}

</script>
  
<style>
body {
    margin: 0;
}

.location-btn {
    border: none;
    font-size: 12px;
}

.custom-control {
    border-radius: 4px;
    display: flex;
    color: #666666;
    background-color: #fff;
    padding: 6px;
    cursor: pointer;
    box-shadow: rgb(0 0 0 / 15%) 1px 2px 1px;
}

.select {
    font-size: 12px;
    color: #666666;
    outline: none;
    border-radius: 6px;
    padding: 1px 10px;
    border: 1px solid var(--vp-custom-block-details-border);
    background-color: var(--vp-custom-block-details-bg);
    appearance: menulist-button;
}
</style>
  