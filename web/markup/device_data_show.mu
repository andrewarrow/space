{{ define "device_data_show" }}
  div my-3
    {{ $item := index . "item" }}
    {{ $item }}
  div id=device-data space-y-3 w-full flex flex-col justify-center items-center
    div
      api
  div space-y-3
    div
      a href=# id=api1
        release_food(bay, grams)
    div
      a href=#
        take_photo
    div
      a href=#
        play_audio(file)
  {{ end }}
