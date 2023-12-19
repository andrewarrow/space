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
        input type=text id=bay placeholder=1 border-2 border-gray-200 p-1 rounded
    div flex
      div w-64
        grams
      div
        input type=text id=grams placeholder=100 border-2 border-gray-200 p-1 rounded
    div flex items-center
      div w-64
        execute
      div text-xs
        div mr-3
          select id=value border-2 border-gray-200 p-1 rounded
            option value=9 
              Now
            option value=10
              Every day 9AM
            option value=10
              Every other day 9AM
            option value=10
              Andrew's iPhone Close
      div
        input type=submit value=Schedule border rounded bg-blue-600 text-white py-2 px-2
  {{ end }}
