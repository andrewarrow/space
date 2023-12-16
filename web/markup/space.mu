div p-3 text-2xl min-h-screen w-full bg-r
  div py-3 text-center
    Welcome to Wifi Router 3000 
  div text-sm text-center
    34 devices
  div space-y-3 w-full flex flex-col justify-center items-center
    {{ template "device" . }}
