div p-3 text-4xl pb-64
  div py-9 text-center
    welcome
  div flex justify-center space-x-9
    div
      button id=login border rounded bg-blue-600 text-white py-2 px-2
        Login
    div
      button id=register whitespace-nowrap border rounded bg-blue-600 text-white py-2 px-2
        Sign Up
  form py-9
    div flex justify-center
      div space-y-9
        div 
          input w-64 type=text id=username autofocus=true required=true placeholder=username
        div
          input w-64 type=password id=password required=true placeholder=password
