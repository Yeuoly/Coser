<template>
    <div class="c-v-gallary-editor px2 py2" ref="galleryRef">
        <div v-if="!gallery.ID">
            <NH3 prefix="bar">
                <NText type="primary">创建</NText>
            </NH3>
            <NResult status="warning">
                请注意，当您创建完正片后，有且只有您当前使用的浏览器可以修改、删除该地点，其他人只拥有查看权限。
            </NResult>
            <NH6 prefix="bar" class="w100">
                <div class="text-13 w100">
                    起个名~
                </div>
                <NInput v-model:value="gallery.name" maxlength="64" show-count />
            </NH6>
            <NH6 prefix="bar" class="w100">
                <div class="text-13 w100">
                    Coser的CN（可以留空）
                </div>
                <NInput v-model:value="gallery.cosers" maxlength="64" show-count />
            </NH6>
            <NH6 prefix="bar" class="w100">
                <div class="text-13 w100">
                    摄影师的CN（可以留空）
                </div>
                <NInput v-model:value="gallery.photographers" maxlength="64" show-count />
            </NH6>
            <NH6 prefix="bar" class="w100">
                <div class="text-13 w100">
                    角色（可以留空）
                </div>
                <NInput v-model:value="gallery.character" maxlength="64" show-count />
            </NH6>
            <NH6 prefix="bar" class="w100">
                <div class="text-13 w100">
                    简单的介绍（可以留空）
                </div>
                <NInput v-model:value="gallery.description" maxlength="1024" type="textarea" show-count />
            </NH6>
            <NH3 prefix="bar" class="w100">
                <NButton block type="primary" @click="createGallery">创建</NButton>
            </NH3>
        </div>
        <div v-else>
            <NGrid cols="2" x-gap="12">
                <NGi>
                    <NH3 prefix="bar">
                        <NText type="primary">修改</NText>
                    </NH3>
                    <NH6 prefix="bar">
                        <div class="text-13 w100">
                            起个名~
                        </div>
                        <NInput v-model:value="gallery.name" maxlength="64" show-count />
                    </NH6>
                    <NH6 prefix="bar">
                        <div class="text-13 w100">
                            人员（可以留空）
                        </div>
                        <NInputGroup>
                            <NInput placeholder="coser" v-model:value="gallery.cosers" maxlength="64" show-count />
                            <NInput placeholder="摄影师" v-model:value="gallery.photographers" maxlength="64" show-count />
                            <NInput placeholder="角色" v-model:value="gallery.character" maxlength="64" show-count />
                        </NInputGroup>
                    </NH6>
                    <NH6 prefix="bar">
                        <div class="text-13 w100">
                            简单的介绍（可以留空）
                        </div>
                        <NInput v-model:value="gallery.description" maxlength="1024" type="textarea" show-count />
                    </NH6>
                    <NH6 prefix="bar">
                        <NUpload multiple directory-dnd :custom-request="uploadImage"
                            v-model:file-list="uploadFileList">
                            <NUploadDragger>
                                <div style="margin-bottom: 12px">
                                    <NIcon size="48" :depth="3">
                                        <ImageIcon />
                                    </NIcon>
                                </div>
                                <NText style="font-size: 16px">
                                    上传新的照片
                                </NText>
                                <NP depth="3" style="margin: 8px 0 0 0">
                                    请确保您上传的图片符合相关国家法律法规，且不要上传敏感数据
                                </NP>
                            </NUploadDragger>
                        </NUpload>
                    </NH6>
                    <NH6 prefix="bar">
                        <div class="text-13 w100">
                            标签（Tag）
                        </div>
                        <NTag closable class="mb1" v-for="tag in gallery.tags" type="info" @close="removeTag(tag)">{{
                            tag.name }}</NTag>
                        <NAutoComplete @select="createTag" :options="tagOptions" placeholder="回车添加" v-model:value="tagName"
                            maxlength="64" show-count @keypress.enter="createTag" />
                    </NH6>
                    <NH3 prefix="bar">
                        <NButton block type="primary" @click="updateGallery">更新</NButton>
                    </NH3>
                </NGi>
                <NGi>
                    <NH3 prefix="bar">
                        <NText type="primary">图片</NText>
                    </NH3>
                    <NGrid cols="1" style="max-height: 80vh; overflow-x: hidden; overflow-y: auto;" y-gap="12">
                        <NGi v-for="image in gallery.images">
                            <div>
                                <CameraEditor v-bind:camera="image.exif"></CameraEditor>
                                <NTag type="error" :bordered="false" size="small" class="clickable"
                                    @click="removeImage(image)">
                                    <template #icon>
                                        <NIcon :component="Trash"></NIcon>
                                    </template>
                                    删除
                                </NTag>
                                <NTag type="info" :bordered="false" size="small" class="clickable"
                                    @click="replaceCamera(image.exif)">
                                    <template #icon>
                                        <NIcon :component="Save"></NIcon>
                                    </template>
                                    将所有照片相机/镜头设置为当前相机/镜头
                                </NTag>
                                <NTag type="primary" :bordered="false" size="small" class="clickable"
                                    @click="updateImage(image)">
                                    >
                                    <template #icon>
                                        <NIcon :component="Save"></NIcon>
                                    </template>
                                    保存更改
                                </NTag>
                            </div>
                            <div class="w100">
                                <NImage lazy :width="imageWidth" :height="300"
                                    :src="image.url + '?x-oss-process=image/resize,mfit,h_800,w_1600'" object-fit="contain"
                                    alt="加载失败"></NImage>
                            </div>
                            <NDivider></NDivider>
                        </NGi>
                    </NGrid>
                </NGi>
            </NGrid>
        </div>
    </div>
</template>

<script setup lang="ts">
import { PropType, VNodeChild, h, onMounted, ref, watch } from 'vue'
import { Gallery, GalleryTag, Image, Place } from '../interface/types'
import { DialogApi, NAutoComplete, NButton, NDivider, NGi, NGrid, NH3, NH6, NIcon, NImage, NInput, NInputGroup, NResult, NTag, NText, NUpload, NUploadDragger, SelectOption, UploadCustomRequestOptions } from 'naive-ui'
import { apiCreateGallery, apiGalleryUploadImage, apiGalleryRemoveImage, apiUpdateGallery, apiGalleryUpdateImage } from '../interface/cos'
import { getBrowserKey } from '../store/key'
import { Camera, Eye, Image as ImageIcon, Save, Telescope, Trash } from '@vicons/ionicons5'
import { getExif } from '../utils/camera'
import { FileInfo } from 'naive-ui/es/upload/src/interface'
import { getDefaultCoser, getDefaultPhotographer, setDefaultCoser, setDefaultPhotographer } from '../store/role'
import { apiCreateTag, apiSearchTag } from '../interface/cos'
import CameraEditor from '../components/CameraEditor.vue'
import { addCommonlyUsedCamera, addCommonlyUsedLens } from '../store/camera'

const imageWidth = ref(0)
const galleryRef = ref<HTMLDivElement | null>(null)
onMounted(() => {
    setTimeout(() => {
        if (!galleryRef.value) {
            return
        }
        imageWidth.value = galleryRef.value.clientWidth / 2 - 12
    })
})

const props = defineProps({
    galleryId: {
        type: Number,
        required: true,
        default: 0
    },
    placeId: {
        type: Number,
        required: true,
        default: 0
    },
    gallery: {
        type: Object as PropType<Gallery>,
        required: false,
        default: () => ({} as Gallery)
    }
})

const emits = defineEmits<{
    (event: 'created', gallery: Gallery): void
    (event: 'updated', gallery: Gallery): void
}>()

const gallery = ref<Gallery>({
    ID: 0,
    CreatedAt: '',
    UpdatedAt: '',
    DeletedAt: '',
    name: '',
    cosers: '',
    photographers: '',
    description: '',
    character: '',
    series: '',
    place_id: 0,
    place: {} as Place,
    tags: [],
    images: [],
    key: '',
})
watch(props.gallery, () => {
    props.gallery.images?.forEach((image, k) => {
        if (!image.exif) {
            console.log(image)
            props.gallery.images[k].exif = {
                camera: image.camera,
                lens: image.lens,
                focal: parseInt(image.focal_length),
                apertureValue: parseFloat(image.aperature),
                exposureTime: parseInt(image.exposure_time),
                iso: parseInt(image.iso),
            }
        }
    })
}, { immediate: true, deep: true })

const replaceCamera = (exif: any) => {
    gallery.value.images.forEach((image, k) => {
        gallery.value.images[k].exif.camera = exif.camera
        gallery.value.images[k].exif.lens = exif.lens
        gallery.value.images[k].exif.focal = exif.focal
    })
    
    // @ts-ignore
    window.$message.success('更新成功，但是您需要手动保存，以避免全自动保存导致的不可挽回的误操作')
}

const creating = ref(false)
const createGallery = async () => {
    if (creating.value) {
        // @ts-ignore
        window.$message.error('创建中')
        return
    }
    creating.value = true
    // @ts-ignore
    const dialog = window.$dialog as DialogApi
    let wait = 0
    if (gallery.value.cosers) {
        const defaultCoser = await getDefaultCoser()
        if (!defaultCoser) {
            wait += 1
            dialog.info({
                title: '缓存确认',
                content: `您是否确认使用 ${gallery.value.cosers} 作为默认的coser？`,
                positiveText: '确认',
                negativeText: '不需要默认coser',
                onPositiveClick: () => {
                    setDefaultCoser(gallery.value.cosers)
                    wait -= 1
                },
                onClose: () => {
                    wait -= 1
                },
                onNegativeClick: () => {
                    setDefaultCoser(' ')
                    wait -= 1
                },
                onMaskClick: () => {
                    wait -= 1
                }
            })
        }
    }
    if (gallery.value.photographers) {
        const defaultPhotographer = await getDefaultPhotographer()
        if (!defaultPhotographer) {
            wait += 1
            dialog.info({
                title: '缓存确认',
                content: `您是否确认使用 ${gallery.value.photographers} 作为默认的摄影师？`,
                positiveText: '确认',
                negativeText: '不需要默认摄影师',
                onPositiveClick: () => {
                    setDefaultPhotographer(gallery.value.photographers)
                    wait -= 1
                },
                onClose: () => {
                    wait -= 1
                },
                onNegativeClick: () => {
                    setDefaultPhotographer(' ')
                    wait -= 1
                },
                onMaskClick: () => {
                    wait -= 1
                }
            })
        }
    }

    while (wait > 0) {
        await new Promise(resolve => setTimeout(resolve, 100))
    }

    const response = await apiCreateGallery(
        props.placeId, gallery.value.name, gallery.value.cosers, gallery.value.description,
        gallery.value.photographers, gallery.value.character, gallery.value.series, gallery.value.tags.map(tag => tag.ID),
        getBrowserKey()
    )

    creating.value = false

    if (!response.isSuccess()) {
        // @ts-ignore
        window.$message.error(response.message)
        return
    }

    gallery.value.ID = response.data?.id || 0

    emits('created', gallery.value)
}

const uploadFileList = ref<FileInfo[]>([])
const uploadImage = async ({
    file,
    onFinish,
    onError,
    onProgress
}: UploadCustomRequestOptions) => {
    if (!file.file) {
        return
    }

    if (['image/jpeg', 'image/jpg', 'image/png', 'image/gif', 'image/bmp', 'image/tiff', 'image/webp', 'image/svg+xml', 'image/x-icon', 'image/heif'].indexOf(file.file.type) === -1) {
        // @ts-ignore
        window.$message.error('不支持的文件类型')
        return
    }

    const exif = await getExif(file.file)

    let index = 0
    let timer = setInterval(() => {
        onProgress({
            percent: index
        })
        index++
        if (index == 99) {
            clearInterval(timer)
            return
        }
    }, 100)

    const response = await apiGalleryUploadImage(
        gallery.value.ID, file.name, getBrowserKey(), file.type || '',
        exif.camera, exif.lens, exif.focal.toString(), exif.apertureValue.toString(), exif.exposureTime.toString(), exif.iso.toString(),
        file.file as File
    )
    clearInterval(timer)
    onFinish()

    uploadFileList.value.pop()

    if (!response.isSuccess()) {
        // @ts-ignore
        window.$message.error(response.message)
        return
    }

    gallery.value.images.push({
        ID: response.data?.id || 0,
        CreatedAt: '',
        UpdatedAt: '',
        DeletedAt: '',
        gallery_id: gallery.value.ID,
        url: response.data?.url || '',
        camera: exif.camera,
        lens: exif.lens,
        focal_length: exif.focal.toString(),
        aperature: exif.apertureValue.toString(),
        exposure_time: exif.exposureTime.toString(),
        iso: exif.iso.toString(),
        exif: exif,
    })

    emits('updated', gallery.value)
}

const updateGallery = async () => {
    const response = await apiUpdateGallery(
        gallery.value.ID, props.placeId, gallery.value.name, gallery.value.cosers, gallery.value.description,
        gallery.value.photographers, gallery.value.character, gallery.value.series, gallery.value.tags.map(tag => tag.ID),
        getBrowserKey()
    )

    if (!response.isSuccess()) {
        // @ts-ignore
        window.$message.error(response.message)
        return
    }

    emits('updated', gallery.value)

    // @ts-ignore
    window.$message.success('更新成功')
}

const tagName = ref('')
let timer = 0
watch(tagName, () => {
    clearTimeout(timer)
    timer = setTimeout(async () => {
        if (!tagName.value) {
            return
        }
        const response = await apiSearchTag(tagName.value)
        if (!response.isSuccess()) {
            // @ts-ignore
            window.$message.error(response.message)
            return
        }

        tagOptions.value = response.data?.tags?.map(item => ({
            value: item.ID.toString(),
            label: item.name
        })) || []
    }, 300)
})
const tagOptions = ref<{
    value: string
    label: string
}[]>()
const removeImage = async (image: Image) => {
    const response = await apiGalleryRemoveImage(gallery.value.ID, image.ID, getBrowserKey())
    if (!response.isSuccess()) {
        // @ts-ignore
        window.$message.error(response.message)
        return
    }

    gallery.value.images = gallery.value.images.filter(item => item.ID !== image.ID)
    emits('updated', gallery.value)
}

const updateImage = async (image: Image) => {
    const response = await apiGalleryUpdateImage(
        gallery.value.ID, image.ID, getBrowserKey(),
        image.exif.camera, image.exif.lens, 
        image.exif.focal.toString(), image.exif.apertureValue.toString(), 
        image.exif.exposureTime.toString(), image.exif.iso.toString(),
    )
    if (!response.isSuccess()) {
        // @ts-ignore
        window.$message.error(response.message)
        return
    }

    // set commonly used cameras and lens
    const camera = image.exif.camera
    const lens = image.exif.lens

    addCommonlyUsedCamera(camera)
    addCommonlyUsedLens(lens)

    gallery.value.images.forEach((item, k) => {
        if (item.ID === image.ID) {
            gallery.value.images[k].camera = image.exif.camera
            gallery.value.images[k].lens = image.exif.lens
            gallery.value.images[k].focal_length = image.exif.focal.toString()
            gallery.value.images[k].aperature = image.exif.apertureValue.toString()
            gallery.value.images[k].exposure_time = image.exif.exposureTime.toString()
        }
    })

    // @ts-ignore
    window.$message.success('更新成功')

    emits('updated', gallery.value)
}

const createTag = async () => {
    setTimeout(async () => {
        const name = tagName.value.trim()
        const response = await apiCreateTag(name)
        tagName.value = ''
        if (!response.isSuccess()) {
            // @ts-ignore
            window.$message.error(response.message)
            return
        }

        gallery.value.tags.push({
            ID: response.data?.id || 0,
            CreatedAt: '',
            UpdatedAt: '',
            DeletedAt: '',
            name: name,
            galleries: []
        })
    })
}

const removeTag = (tag: GalleryTag) => {
    gallery.value.tags.splice(gallery.value.tags.indexOf(tag), 1)
}

onMounted(async () => {
    if (!props.placeId) {
        // @ts-ignore
        window.$message.error('Place id is required')
        return
    }
    if (props.galleryId) {
        gallery.value.ID = props.galleryId
        gallery.value.name = props.gallery.name
        gallery.value.cosers = props.gallery.cosers
        gallery.value.photographers = props.gallery.photographers
        gallery.value.description = props.gallery.description
        gallery.value.character = props.gallery.character
        gallery.value.series = props.gallery.series
        gallery.value.tags = props.gallery.tags
        gallery.value.images = props.gallery.images
        return
    } else {
        const photographers = await getDefaultPhotographer()
        gallery.value.photographers = photographers
        const coser = await getDefaultCoser()
        gallery.value.cosers = coser
    }
})

</script>

<style scoped></style>