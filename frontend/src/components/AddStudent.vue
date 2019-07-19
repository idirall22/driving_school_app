<template>
  <div id="studentDetails">
    <div class="row">
      <Header initTitle="Ajouter Un Étudiant"></Header>

      <!-- Message  -->
       <Message
        v-on:close-alert="closeAlert"
        :create.sync="studentCreated"
        :errorCreate.sync="errorCreateStudent"
        :subject="subjectMessage">
      </Message>

    </div>

    <!-- Errors fields -->
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
            <label for="firstName">Prénom</label>
            <input v-model="firstName" type="text" class="form-control"
            id="firstName" placeholder="" value="" required >
          </div>

          <!-- first name -->
          <div class="col-md-6 mb-3">
            <label class="d-flex justify-content-end" for="arfirstName">الإسم</label>
            <input v-model="arFirstName"
              dir="rtl" type="text" class="form-control"
            id="arfirstName" placeholder="" value="" required >
          </div>

          <!-- last name -->
          <div class="col-md-6 mb-3">
            <label for="lastName">Nom De Famille</label>
            <input v-model="lastName" type="text" class="form-control"
            id="lastName" placeholder="" value="" required >
          </div>

          <!-- last name -->
          <div class="col-md-6 mb-3">
            <label class="d-flex justify-content-end" for="lastName">اللقب</label>
            <input v-model="arlastName" type="text" class="form-control"
            id="lastName" placeholder="" value="" required >
          </div>
        </div>

        <div class="row">

          <!-- maiden name -->
          <div class="col mb-3">
            <label for="address">Nom De Jeune Fille</label>
            <input v-model="maidenName" type="text" class="form-control" id="maidenName"
            placeholder="" required >
          </div>

          <!-- maiden name -->
          <div class="col mb-3">
            <label class="d-flex justify-content-end" for="address">إسم العائلة قبل الزواج</label>
            <input v-model="maidenName" type="text" class="form-control" id="maidenName"
            placeholder="" required >
          </div>

        </div>

        <div class="row">

          <!-- birthday -->
          <div class="col-4 mb-3">
            <label for="birthday">Date de Naissance</label>
            <vue-bootstrap-datetimepicker v-model="birthday" :config="options"
            required></vue-bootstrap-datetimepicker>
          </div>

          <!-- birth city -->
          <div class="col-4 mb-3">
            <label for="birthCity">Wilaya de Naissance</label>
            <input v-model="birthCity" type="text" class="form-control"
              id="birthCity" placeholder="Oran" required >
          </div>

          <!-- Country -->
          <div class="col mb-3">
            <label for="country">Pays De Naissance</label>
            <input v-model="country" type="text" class="form-control"
            id="country" placeholder="Algerie" required >
          </div>

        <!-- gender -->
        <div class="col-4 mb-3">
          <label for="gender">Sexe</label>
          <select v-model="gender" class="custom-select d-block w-100"
            id="gender" required>

            <option value="man">Man</option>
            <option value="woman">Woman</option>
          </select>
        </div>

        <!-- registred date -->
        <div class="col-4 mb-3">
          <label for="registredDate">Date d'Inscription</label>
          <vue-bootstrap-datetimepicker v-model="registredDate" :config="options"
          ></vue-bootstrap-datetimepicker>
        </div>

        <!-- phone number -->
        <div class="col mb-3">
          <label for="address">Numero De Telephone</label>
          <input v-model="phoneNumber" type="text" class="form-control"
          id="address" placeholder="XX-XX-XX-XX-XX" required >
        </div>

      </div>

      <div class="row">
        <!-- address -->
        <div class="col-12 mb-3">
          <label for="address">Adresse</label>
          <input v-model="addressStreet" type="text" class="form-control" id="address" placeholder="1234 Main St"
            required >
        </div>

        <!-- department -->
        <div class="col-4 mb-3">
          <label for="department">Daira</label>
          <input v-model="department" type="text" class="form-control" id="address" placeholder="Bir el djir"
            required >
        </div>

        <!-- city -->
        <div class="col-4 mb-3">
          <label for="city">Wilaya</label>
          <input v-model="city" type="text" class="form-control" id="city"
            placeholder="Oran" required >
        </div>
      </div>

      <div class="row">

        <!-- fileNumber -->
        <div class="col-4 mb-3">
          <label for="fileNumber">Numero De Dossier</label>
          <input v-model="fileNumber" type="text" class="form-control" id="fileNumber" placeholder="" required
            >
        </div>

        <!-- job -->
        <div class="col-4 mb-3">
          <label for="job">Travaille</label>
          <input v-model="job" type="text" class="form-control" id="job" placeholder=""
            required >
        </div>

      </div>
      <!-- button add student -->
      <button @click.prevent="addStudent()"
        class="btn btn-primary btn-lg btn-block" type="submit">Ajouter L'Étudiant</button>
    </form>
  </div>
</template>

<script>
import VueBootstrapDatetimepicker from 'vue-bootstrap-datetimepicker';
import Header from './parts/Header';
import Message from './parts/Message';
import moment from 'moment';

export default {
  name: "addStudent",
  components: {
    VueBootstrapDatetimepicker,
    Header,
    Message
  },
  data: () => ({
    options:{
      format:"YYYY-MM-DD",
      useCurrent: false
    },
    arabic:"",
    arFirstName:"",
    subjectMessage:"Student Created",
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
        if(this.maidenName == ""){
          this.maidenName = "None";
        }
        this.student = {
          "first_name": this.firstName,
          "last_name": this.lastName,
          "maiden_name": this.maidenName,
          "phone_number": this.phoneNumber,
          "birthday": moment(this.birthday).format(),
          "birth_city": this.birthCity,
          "job": this.job,
          "city": this.city,
          "file_number": this.fileNumber,
          "gender": this.gender,
          "country": this.country,
          "address_street": this.addressStreet,
          "department": this.department,
          "registred_date": moment(this.registredDate).format(),
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
    },
  }
}
"أ"
</script>
<style scoped>
.ar{
  direction: rtl;
}
</style>