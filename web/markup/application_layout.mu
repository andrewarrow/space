html attr-1
  head
    title
      {{ index . "title" }}
    script src=https://cdn.tailwindcss.com
    link rel=stylesheet type=text/css href=/assets/css/tail.min.css
  body
    div id=flash bg-red-600 text-white text-center
      {{ index . "flash" }}
    div bg-r overflow-x-auto pl-20 pr-20
      {{ index . "content" }}
