// @ts-ignore
import piexif from 'piexifjs'

type EXIF = {
    camera: string
    lens: string
    focal: string
    apertureValue: string
    exposureTime: string
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

            let focal = ''
            if (_focal) {
                focal = '' + _focal[0] / _focal[1]
            }

            let apertureValue = ''
            if (_apertureValue) {
                apertureValue = '' + Math.pow(2, _apertureValue[0] / _apertureValue[1] / 2).toFixed(1)
            }

            let exposureTime = ''
            if (_exposureTime) {
                exposureTime = '' + `${_exposureTime[0]}/${_exposureTime[1]}`
            }

            resolve({
                camera,
                lens,
                focal,
                apertureValue,
                exposureTime
            })
        } catch (e) {
            resolve({
                camera: '',
                lens: '',
                focal: '',
                apertureValue: '',
                exposureTime: ''
            })
        }
    }
})