// @ts-ignore
import piexif from 'piexifjs'

export type EXIF = {
    camera: string
    lens: string
    focal: number
    apertureValue: number
    exposureTime: number
    iso: number
}

export const getExif = (file: File) => new Promise<EXIF>((resolve) => {
    const reader = new FileReader()
    reader.readAsDataURL(file as File)
    reader.onload = function (e) {
        // read exif
        const imageData = e.target?.result
        try {
            const exifObj = piexif.load(imageData as string)
            const _camera = exifObj['0th'][piexif.ImageIFD.Make]
            const _model = exifObj['0th'][piexif.ImageIFD.Model]
            const _lens = exifObj['Exif'][piexif.ExifIFD.LensMake].trim() + ' ' + exifObj['Exif'][piexif.ExifIFD.LensModel].trim()
            const _focal = exifObj['Exif'][piexif.ExifIFD.FocalLength]
            const _apertureValue = exifObj['Exif'][piexif.ExifIFD.ApertureValue]
            const _maxApertureValue = exifObj['Exif'][piexif.ExifIFD.MaxApertureValue]
            const _exposureTime = exifObj['Exif'][piexif.ExifIFD.ExposureTime]
            const _iso = exifObj['Exif'][piexif.ExifIFD.ISOSpeedRatings]

            let camera = ''
            if (_camera) {
                camera = _camera
                if (_model) {
                    camera += ' ' + _model
                }
            }

            let lens = ''
            if (_lens) {
                // @ts-ignore
                lens = _lens.replaceAll('\x00', '').replaceAll('\u0000', '')
            }

            let focal = 0
            if (_focal) {
                focal = _focal[0] / _focal[1]
            }

            let apertureValue = 0
            if (_apertureValue) {
                apertureValue = Math.pow(2, _apertureValue[0] / _apertureValue[1] / 2)
                apertureValue = Math.round(apertureValue * 10) / 10
            }

            let exposureTime = 0
            if (_exposureTime) {
                exposureTime = _exposureTime[1] / _exposureTime[0]
            }

            let iso = 0
            if (_iso) {
                iso = _iso
            }

            resolve({
                camera,
                lens,
                focal,
                apertureValue,
                exposureTime,
                iso
            })
        } catch (e) {
            resolve({
                camera: '',
                lens: '',
                focal: 0,
                apertureValue: 0,
                exposureTime: 0,
                iso: 0
            })
        }
    }
})