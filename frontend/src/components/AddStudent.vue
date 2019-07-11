<template>
  <div id="studentDetails">
    <div class="row">
      <div class="col mb-5 mt-3">
        <h1>Add Student</h1>
      </div>
      {{errorCreateStudent}}
      <div v-if="studentCreated"
        class="col mb-5 mt-3 mr-3 alert alert-success" role="alert">
        <button @click.prevent="closeAlert()" type="button" class="close"
            data-dismiss="alert" aria-label="Close">
           <span aria-hidden="true">&times;</span>
         </button>
        <h4 class="alert-heading">Student Created!</h4>
        <hr>
        <p>The student was added Successfuly</p>
      </div>
      <div v-if="!studentCreated && errorCreateStudent != null" class="col mb-5 mt-3 mr-3 alert alert-warning" role="alert">
        <button @click.prevent="closeAlert()" type="button" class="close"
            data-dismiss="alert" aria-label="Close">
           <span aria-hidden="true">&times;</span>
         </button>
        <h4 class="alert-heading">Could not create a new student!</h4>
        <hr>
        <p>The student was not created</p>
      </div>
    </div>

    <div v-if="errors.length > 0" class="col mb-5 mt-3 mr-3 alert alert-danger" role="alert">
      <h4 class="alert-heading">there are some Errors!</h4>
      <hr>
        <p v-for="err in errors" :key="err">
          Error with "{{err}}"</p>
    </div>


    <!-- form -->
      <form class="needs-validation" novalidate="">
        <div class="row">

          <!-- first name -->
          <div class="col-md-6 mb-3">
            <label for="firstName">First name</label>
            <input v-model="firstName" type="text" class="form-control"
            id="firstName" placeholder="" value="" required >
          </div>

          <!-- last name -->
          <div class="col-md-6 mb-3">
            <label for="lastName">Last name</label>
            <input v-model="lastName" type="text" class="form-control"
            id="lastName" placeholder="" value="" required >
          </div>
        </div>

        <div class="row">

          <!-- maiden name -->
          <div class="col mb-3">
            <label for="address">Maiden Name</label>
            <input v-model="maidenName" type="text" class="form-control" id="maidenName"
            placeholder="" required >
          </div>

          <!-- phone number -->
          <div class="col mb-3">
            <label for="address">Phone Number</label>
            <input v-model="phoneNumber" type="text" class="form-control"
            id="address" placeholder="" required >
          </div>

        </div>

        <div class="row">

          <!-- birthday -->
          <div class="col-4 mb-3">
            <label for="birthday">Birthday</label>
            <vue-bootstrap-datetimepicker v-model="birthday" :config="options"
            required></vue-bootstrap-datetimepicker>
          </div>

          <!-- birth city -->
          <div class="col-4 mb-3">
            <label for="birthCity">BirthCity</label>
            <select v-model="birthCity" class="custom-select d-block" id="birthCity"
              required>
              <option value="oran">Oran</option>
            </select>
          </div>

          <!-- Country -->
          <div class="col mb-3">
            <label for="country">Country</label>
            <input v-model="country" type="text" class="form-control"
            id="country" placeholder="" required >
          </div>

        <!-- gender -->
        <div class="col-4 mb-3">
          <label for="gender">Gender</label>
          <select v-model="gender" class="custom-select d-block w-100"
            id="gender" required>

            <option value="man">Man</option>
            <option value="woman">Woman</option>
          </select>
        </div>

        <!-- registred date -->
        <div class="col-4 mb-3">
          <label for="registredDate">Registred Date</label>
          <vue-bootstrap-datetimepicker v-model="registredDate" :config="options"
          ></vue-bootstrap-datetimepicker>
        </div>

      </div>

      <div class="row">
        <!-- address -->
        <div class="col-12 mb-3">
          <label for="address">Address</label>
          <input v-model="addressStreet" type="text" class="form-control" id="address" placeholder="1234 Main St"
            required >
        </div>

        <!-- department -->
        <div class="col-4 mb-3">
          <label for="department">Department</label>
          <input v-model="department" type="text" class="form-control" id="address" placeholder="Bir el djir"
            required >
        </div>

        <!-- city -->
        <div class="col-4 mb-3">
          <label for="city">City</label>
          <select v-model="city" class="custom-select d-block" id="city"
            required>
            <option value="oran">Oran</option>
          </select>
        </div>
      </div>

      <div class="row">

        <!-- fileNumber -->
        <div class="col-4 mb-3">
          <label for="fileNumber">FileNumber</label>
          <input v-model="fileNumber" type="text" class="form-control" id="fileNumber" placeholder="" required
            >
        </div>

        <!-- job -->
        <div class="col-4 mb-3">
          <label for="job">Job</label>
          <input v-model="job" type="text" class="form-control" id="job" placeholder=""
            required >
        </div>

      </div>
      <!-- button add student -->
      <button @click.prevent="addStudent()"
        class="btn btn-primary btn-lg btn-block" type="submit">Add Student</button>
    </form>
  </div>
</template>

<script>
import VueBootstrapDatetimepicker from 'vue-bootstrap-datetimepicker';

export default {
  name: "addStudent",
  components: {
    VueBootstrapDatetimepicker
  },
  data: () => ({
    options:{
      // format:"YYYY/MM/DD",
      useCurrent: false
    },
    student:{},

    fileNumber: "",
    firstName: "",
    lastName: "",
    maidenName: "",
    phoneNumber: "",
    birthday: "",
    birthCity: "",
    job: "",
    gender: "",
    country: "",
    city: "",
    department: "",
    addressStreet: "",
    registredDate: "",

    studentCreated: false,
    errorCreateStudent: null,
    errors: [],
  }),
  methods:{
    addStudent:function(){
      this.errors = []
      if (this.firstName == ""){
        this.errors.push("First Name")
      }
      if (this.lastName == ""){
        this.errors.push("Last Name")
      }
      if (this.phoneNumber == ""){
        this.errors.push("Phone Number")
      }
      if (this.birthday == ""){
        this.errors.push("Birthday")
      }
      if (this.birthCity == ""){
        this.errors.push("Birth City")
      }
      if (this.country == ""){
        this.errors.push("Country")
      }
      if (this.job == ""){
        this.errors.push("Job")
      }
      if (this.fileNumber == ""){
        this.errors.push("File Number")
      }
      if (this.city == ""){
        this.errors.push("City")
      }
      if (this.department == ""){
        this.errors.push("Department")
      }
      if (this.addressStreet == ""){
        this.errors.push("Address")
      }
      if (this.registredDate == ""){
        this.errors.push("Registred Date")
      }
      if(this.errors.length == 0){
        this.student = {
          "first_name": this.firstName,
          "last_name": this.lastName,
          "maiden_name": this.maidenName,
          "phone_number": this.phoneNumber,
          "birthday": this.birthday,
          "birth_city": this.birthCity,
          "job": this.job,
          "city": this.city,
          "file_number": this.fileNumber,
          "gender": this.gender,
          "country": this.country,
          "department": this.department,
          "address_street": this.addressStreet,
          "registred_date": this.registredDate,
        }
          window.backend.Service.CreateStudent(this.student).
          then(
            data=>{this.data = data;},
            err=>{this.errorCreateStudent = err;},
          )
          if(this.errorCreateStudent != null){
            this.studentCreated = false;
          }else{
            this.studentCreated = true;
        }
      }
    },
    closeAlert:function(){
      this.studentCreated = false;
      this.errorCreateStudent = null;
    }
  }
}

</script>
<style scoped>
</style>