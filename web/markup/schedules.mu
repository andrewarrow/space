{{ define "schedules" }}
  div space-y-3 flex flex-col justify-center items-center
    form w-7/8 p-3 space-y-3 id=add-schedule-form bg-white rounded
      div flex space-x-3
        div
          Every day at
        div
          select id=value text-black
            option value=9 
              9AM
            option value=10
              10AM
      div flex space-x-3 items-center
        div
          input type=text w-full id=name placeholder=name border-2 border-grey-400 p-1 rounded
        div text-xs
          input type=submit value=Add border rounded bg-blue-600 text-white py-2 px-2
    div id=schedule-list w-full
      {{ . }}
  {{ end }}
