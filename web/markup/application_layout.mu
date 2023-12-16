html attr-1
  head
    {{ if index . "USE_LIVE_TEMPLATES" }}
      script src=https://cdn.tailwindcss.com
    {{ else }}
      {{ $build := index . "build" }}
      link rel=stylesheet type=text/css href=/assets/css/tail.min.css?id!{{$build}}
    {{ end }}
    script src=/assets/javascript/wasm_exec.js
    title
      {{ index . "title" }}
  body
    div id=flash bg-red-600 text-white text-center fixed top-0 left-0 w-full
      {{ index . "flash" }}
    div bg-blue-100 overflow-x-auto pl-3 pr-3 w-full
      {{ index . "content" }}
    {{ index . "wasm" }}
