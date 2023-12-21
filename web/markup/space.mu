div p-0 w-full flex justify-center items-center
  div p-3 w-full md:w-1/2 
    div py-3 flex items-center justify-center
      div pl-1 text-sm
        a href=# id=schedules
          schedules
      div w-full text-center 
        homeOS
      div ml-auto pr-1 text-sm
        a href=# id=messages
          rooms
    div text-sm text-center
      span id=total
        0
      span
        devices
    div id=devices space-y-3 w-full flex flex-col justify-center items-center
      &nbsp;
  div fixed top-0 left-1/2 transform -translate-x-1/2 h-full hidden id=modal w-full md:w-1/2 overflow-y-auto bg-red-900 rounded
    div flex justify-left items-center p-6 
      div 
        a href=# id=back
          [BACK]
      div text-center w-full
        Title is longer
    div id=modal-content p-3 h-full bg-white
      modal
