html attr-1
  head
    script src=https://cdn.tailwindcss.com
    title
      {{ index . "title" }}
  body
    div id=flash bg-red-600 text-white text-center
      {{ index . "flash" }}
    div bg-r overflow-x-auto pl-20 pr-20
      {{ index . "content" }}
