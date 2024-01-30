document.addEventListener('alpine:init', async() => {
   Alpine.store("data", ({
      all_files: [],
      col_structure: [],
      col_options: [],
      form_data: {},
      form_result: [],

      async load(){
         await load_col_structure();
         await load_col_options();
         await load_files();
      },
      update_form(key_name, value){
         if (!(key_name in this.form_data) ) {
            this.form_data[key_name] = []
            this.form_data[key_name].push(value)   
         } else if (key_name in this.form_data) {
            if (this.form_data[key_name].includes(value)){
               this.form_data[key_name].pop(value)
               return
            }
            this.form_data[key_name].push(value)   
         }
      },
      calculate_form_result(){
         this.form_result=JSON.stringify(this.form_data)
         
         col_structure = JSON.parse(JSON.stringify(this.col_structure))
         form_data = JSON.parse(JSON.stringify(this.form_data))
         all_files = JSON.parse(JSON.stringify(this.all_files))

         result = []     
         toPop = []
         form_data = updateKeys(form_data, col_structure)
         
         for (const key in form_data) {
            const values = form_data[key];
            for (const file of all_files) {
                if (Array.isArray(file[key])) {
                    let good = true;
                    for (const value of values) {
                        if (!file[key].includes(value)) {
                            good = false;
                            break;
                        }
                    }
                    if (good && !result.includes(file)) {
                        result.push(file);
                    } else if (!good && !toPop.includes(file)) {
                        toPop.push(file);
                    }
                }
            }
        }
        
        for (const key in form_data) {
            const values = form_data[key];
            for (const file of all_files) {
                if (!Array.isArray(file[key])) {
                    for (const value of values) {
                        if (value === file[key]) {
                            if (!result.includes(file)) {
                                result.push(file);
                            }
                        } else if (!toPop.includes(file)) {
                            toPop.push(file);
                        }
                    }
                }
            }
        }
        
        for (const res of result) {
            for (const willPop of toPop) {
                const index = result.indexOf(willPop);
                if (index !== -1) {
                    result.splice(index, 1);
                }
            }
        }
        
         this.form_result = result
      }   
   }))
})

function doesObjectExist(list, targetObject) {
   for (let i = 0; i < list.length; i++) {
      if (objectsAreEqual(list[i], targetObject)) {
         return true;
      }
   }
   return false; 
}
 
function objectsAreEqual(obj1, obj2) {
   const keys1 = Object.keys(obj1);
   const keys2 = Object.keys(obj2);

   if (keys1.length !== keys2.length) {
      return false;
   }

   for (let key of keys1) {
      if (obj1[key] !== obj2[key]) {
         return false;
      }
   }

   return true;
}

function updateKeys(formData, colStructure) {
   const updatedFormData = {};
 
   for (const [formKey, formValue] of Object.entries(formData)) {
     const obj = colStructure[0]
     const colKey = Object.keys(obj).find(
       key => obj[key] === formKey
     );

 
     if (colKey) {
       updatedFormData[colKey] = formValue;
     } else {
       updatedFormData[formKey] = formValue;
     }
   }
 
   return updatedFormData;
}

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
      strucure_data.push(
         {
            name: col_names[key],
            options: data[key]
         }
      )
   }
   
   Alpine.store("data").col_options = strucure_data
}

