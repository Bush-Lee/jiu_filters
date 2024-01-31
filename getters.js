async function getColOptions() {
   var myHeaders = new Headers();

   var requestOptions = {
      method: 'GET',
      headers: myHeaders,
      redirect: 'follow'
   }

   var url = "https://bush-lee.github.io/jiu_filters/static/col_options.json"
   // var url = "localhost:8081/col_options.json"
   const response = await fetch(url, requestOptions)
   const data = await response.json()
   
   return data
}

async function getColDescription() {
   var myHeaders = new Headers();

   var requestOptions = {
      method: 'GET',
      headers: myHeaders,
      redirect: 'follow'
   }

   var url = "https://bush-lee.github.io/jiu_filters/static/col_description.json"
   // var url = "localhost:8081/col_description.json"
   const response = await fetch(url, requestOptions)
   const data = await response.json()
   
   return data
}

async function getAllFiles() {
   var myHeaders = new Headers();

   var requestOptions = {
      method: 'GET',
      headers: myHeaders,
      redirect: 'follow'
   }

   var url = "https://bush-lee.github.io/jiu_filters/static/all_files.json"
   // var url = "localhost:8081/all_files.json"
   const response = await fetch(url, requestOptions)
   const data = await response.json()
   
   return data
}

