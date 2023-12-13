div p-3 text-2xl h-screen
  div py-3 text-center
    dr. jones
  div space-y-3
    div flex justify-center
      div
        button id=profile border rounded bg-blue-600 text-white py-2 px-2
          Profile
    div flex justify-center
      div
        button whitespace-nowrap id=create border rounded bg-blue-600 text-white py-2 px-2
          Create Ticket
    div flex justify-center
      div mt-3
        a href=/tickets whitespace-nowrap id=view border rounded bg-blue-600 text-white py-2 px-2
          View Tickets
  div bg-r fixed top-0 left-0 w-full h-screen hidden id=modal
    div text-right mr-3
      a href=# id=x
        X
    form py-9 id=ticket-form method=POST
      div flex justify-center
        div space-y-3
          div 
            input w-64 type=text id=title required=true placeholder=title
          div
            description
          div
            textarea w-64 h-64 id=description placeholder=description
          div
            input w-64 type=submit value=Create border rounded bg-blue-600 text-white py-2 px-2
