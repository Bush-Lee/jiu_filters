document.addEventListener('alpine:init', async() => {
   Alpine.store("data", ({
      // Store data
      all_files: [],
      col_structure: [],
      col_options: [],

      // Form
      form_result: [],

      // Page data
      current_option_index: 0,
      current_option_select: "",

      async load(){
         await load_col_structure();

         await load_col_options();
         this.current_option_select = this.col_options[0];

         await load_files();
      },        

      // Form options 
      reset_form() {
         for (const [index, element] of this.col_options.entries()) {
            for (const [i, el] of element.options.entries()) {
               if (el.checked) {
                  el.checked = false;
               }
            }
         }

         this.current_option_index= 0,
         this.current_option_select = this.col_options[0];
      },

      // Result calculator
      get_form_options() {
         checked_options = new Map(); 
         for (const [index, element] of this.col_options.entries()) {
            key = find_key_by_value(this.col_structure[0], element.name)
            if (!checked_options.has(element.name)) {
               checked_options[key] = []
            }
            for (const [i, el] of element.options.entries()) {
               if (el.checked) {
                  checked_options[key].push(el.value);
               }
            }
         }
         this.calculate_form_results(checked_options);
      },

      calculate_form_results(checked_options) {
         result = new Array();
         for (const file of this.all_files) {
            // Tot rahatul asta de if ii inutil, dar mi-e lene sa si verific
            // ar trebui sa ai doar pushul 
            result_index = find_index_by_key(result, file)
            if (result_index == -1) {
               result.push(
                  {
                     file: file,
                     i: 0
                  }
               )
               
               result_index = result.length - 1
            }

            for (let key in file) {
               if (key == "Link") {
                  continue
               }

               if (Array.isArray(file[key])) {
                  for (let value in file[key]) {
                     if (checked_options[key].includes(value)) {
                        result[result_index].i += 1
                     }
                  }
               } else {
                  if (checked_options[key].includes(file[key])) {
                     result[result_index].i += 1
                  }
               }
            }  
         }
         
         result.sort((a, b) => a.i - b.i);
         result = result.slice(-5)

         file_result = new Array();
         for (let res of result) {
            file_result.push(res.file)
         }
         this.form_result = file_result
      },

      // Page option parser
      next_option() {
         this.current_option_index += 1;
         // TODO: Handle special cases
         this.current_option_select = this.col_options[this.current_option_index];
      },

      previous_option() {
         // TODO: Handle special cases
         this.current_option_index -= 1;
         this.current_option_select = this.col_options[this.current_option_index];
      },
   }))
})

function find_index_by_key(list, key) {
   for (let i = 0; i < list.length; i++) {
      if (list[i].key === key) {
        return i;
      }
    }
    return -1; // Return -1 if the key is not found in any object
}

function find_key_by_value(obj, x) {
   for (const key in obj) {
     if (obj[key] === x) {
       return key;
     }
   }
   return null;
 };

async function load_col_structure() {
   col_structure = await getColDescription()
   col_structure = JSON.parse(JSON.stringify(col_structure))
   Alpine.store("data").col_structure = col_structure
}

async function load_files() {
   final_data = await getAllFiles()
   final_data = JSON.parse(JSON.stringify(final_data))
   Alpine.store("data").all_files = final_data
}

async function load_col_options() {
   data = await getColOptions()
   data = JSON.parse(JSON.stringify(data))
   
   strucure_data = []
   col_names = Alpine.store("data").col_structure
   col_names = JSON.parse(JSON.stringify(col_names))
   col_names = col_names[0]

   for (key in col_names) {
      // TODO: Handle Link properly
      if (key == "Link") {
         continue;
      }

      for (const [index, element] of data[key].entries()) {
         data[key][index] = {
            value: element,
            checked: false
         }
      }

      strucure_data.push(
         {
            name: col_names[key],
            options: data[key],
         }
      )
   }
   Alpine.store("data").col_options = strucure_data
}