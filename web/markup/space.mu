div p-0 h-full overflow-y-auto bg-r w-full flex justify-center items-center
  div p-3 text-2xl min-h-screen w-full md:w-1/2 bg-r
    div py-3 text-center
      homeOS
    div text-sm text-center
      34 devices
    div id=devices space-y-3 w-full flex flex-col justify-center items-center
      {{ $items := index . "items" }}
      {{ range $i, $item := $items }}
      {{ template "cat" $item }}
      {{ end }}
  div fixed top-0 left-1/2 transform -translate-x-1/2 h-full hidden id=modal w-full md:w-1/2 overflow-y-auto
    div flex justify-left items-center p-6 w-full bg-blue-300
      div w-full bg-red-600
        a href=# id=x
          [BACK]
    div id=modal-content p-3 h-full  overflow-y-auto bg-white
      modal
