<div align="center">

<h1>      videoto     </h1>
    <small>hls, ts, mp4 to ts, video to screenshots, golang</small>
</div>
 
## About 
Golang библиотека для конвертации видео, на основе ffmpeg.

В библиотеке минимум того что нужно для приличного видеохостинга. Генерит webp превью, скриншоты и разделение на ts файлы. По типу Рутьюба или ВК, они как раз ts используют.

   

## Install

    go get slavaWins/videoto_ts_forhosting_go


## Req ffmpeg

Используется ffmpeg и ffprobe 

https://www.ffmpeg.org/


В докер файле для alpine можно использовать 

    RUN apk add  --no-cache ffmpeg  
    RUN apk add  --no-cache ffprobe  


## Example
Создаем папки testfile и output.  В testfile загружаем видеофайл

    videoto.Screenshots("testfile/input.mp4", "output", 3)
    videoto.TsSegmentation("testfile/input.mp4", "output", "segment")
    videoto.PreviewWebp("testfile/input.mp4", "output")
    videoto.PreviewMp4("testfile/input.mp4", "output")

На выходе получаться 
- 3 скрина
- плейлист m3u8 и сегменты ts 
- и превьюшка в webp формате на 5 секунд, сжатая

## Windows local testing

Скачиваем 
https://www.ffmpeg.org/download.html

Распаковываем в какую-то папку. Добавлям эту папку в Path в переменные среды. Что бы из консоли была доступна функция ffmpeg.