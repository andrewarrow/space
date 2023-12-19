{{ define "device_data_invoke" }}
  div my-3
    {{ $item := index . "item" }}
    {{ $item }}
  div id=device-data space-y-3 w-full flex flex-col justify-center items-center
    div
      release_food(bay, grams)
  form space-y-3 mt-3
    div flex
      div w-64
        bay
      div
        input type=text id=bay placeholder=1 border-2 border-grey-400 p-1 rounded
    div flex
      div w-64
        grams
      div
        input type=text id=grams placeholder=100 border-2 border-grey-400 p-1 rounded
    div flex
      div w-64
        execute
      div text-xs
        input type=submit value=Now border rounded bg-blue-600 text-white py-2 px-2
  {{ end }}
