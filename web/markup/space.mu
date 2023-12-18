div p-0 h-full overflow-y-auto w-full flex justify-center items-center
  div p-3 text-2xl min-h-screen w-full md:w-1/2 
    div py-3 flex items-center justify-center
      div pl-1 text-sm
        schedules
      div w-full text-center 
        homeOS
      div ml-auto pr-1 text-sm
        messages 
    div text-sm text-center
      span id=total
        0
      span
        devices
    div id=cats space-y-3 w-full flex flex-col justify-center items-center
      {{ $items := index . "items" }}
      {{ range $i, $item := $items }}
      {{ template "cat" $item }}
      {{ end }}
  div fixed top-0 left-1/2 transform -translate-x-1/2 h-full hidden id=modal w-full md:w-1/2 overflow-y-auto
    div flex justify-left items-center p-6 w-full bg-blue-300
      div w-full 
        a href=# id=back
          [BACK]
    div id=modal-content p-3 h-full  overflow-y-auto bg-white
      modal
