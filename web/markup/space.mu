div p-3 text-2xl min-h-screen w-full bg-r
  div py-3 text-center
    Welcome to Wifi Router 3000 
  div text-sm text-center
    34 devices
  div id=devices space-y-3 w-full flex flex-col justify-center items-center
    {{ $items := index . "items" }}
    {{ range $i, $item := $items }}
    {{ template "device" $item }}
    {{ end }}
  div bg-r fixed top-0 left-0 w-full min-h-screen hidden id=modal
    div text-right mr-3
      a href=# id=x
        X
    div id=modal-content p-3
      modal
