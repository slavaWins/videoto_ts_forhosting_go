<div align="center">

<h1> videoto_ts_forhosting_go</h1>
 
</div>
 
## About 

Библиотека для видеохостинга. Генерит превью, скриншоты и разделение на ts файлы

   

## Install

    go get slavaWins/videoto_ts_forhosting_go


## Req ffmpeg

Используется ffmpeg и ffprobe 

https://www.ffmpeg.org/


В докер файле для alpine можно использовать 

    RUN apk add  --no-cache ffmpeg  
    RUN apk add  --no-cache ffprobe  


## Use
Создаем папки testfile и output  в testfile загружаем видеофайл

    videoto_ts_forhosting_go.Screenshots("testfile/input.mp4", "output", 3)
    videoto_ts_forhosting_go.TsSegmentation("testfile/input.mp4", "output", "segment")
    videoto_ts_forhosting_go.WebpPreview("testfile/input.mp4", "output")

На выходе получаться 
- 3 скрина
- плейлист m3u8 и сегменты ts 
- и превьюшка в webp формате на 5 секунд, сжатая 