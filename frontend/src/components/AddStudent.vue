<template>
  <div id="studentDetails">
    <b-alert
      v-model="studentCreated"
      :variant="alertVariant"
      dismissible
      @dismissed="closeAlert()">
      {{alertMessage}}
    </b-alert>

    <Header initTitle="Ajouter Un Étudiant"></Header>

    <b-form @submit="addStudent">
      <div class="row">
          <div class="col">
            <!-- First name  -->
            <b-form-group
              id="firstName-group"
              label="Prénom:"
              label-for="firstName">

            <b-form-input
              id="firstName"
              v-model="firstName"
              type="text"
              placeholder="Prénom"
              required
            ></b-form-input>
            </b-form-group>
            <!-- end First name  -->

            <!-- Last name  -->
            <b-form-group
              id="lastName-group"
              label="Nom:"
              label-for="lastName">

              <b-form-input
                id="lastName"
                v-model="lastName"
                type="text"
                placeholder="Nom"
                required
              ></b-form-input>
            </b-form-group>
            <!-- end Last name  -->

            <!-- Maiden name  -->
            <b-form-group
              id="maidenName-group"
              label="Nom De Jeune Fille:"
              label-for="maidenName">

              <b-form-input
                id="maidenName"
                v-model="maidenName"
                type="text"
                placeholder= "Nom De Jeune Fille"
                required
              ></b-form-input>
            </b-form-group>
            <!-- end Maiden name  -->

          </div>

          <div class="col">
            <!-- AraFirst name  -->
            <b-form-group
              id="firstName-group"
              label=":الإسم"
              label-for="firstName">

              <b-form-input
              id="firstName"
                v-model="arFirstName"
                type="text"
                placeholder="الإسم"
                required
              ></b-form-input>
            </b-form-group>
            <!-- end First name  -->

            <!-- Arabic Last name  -->
            <b-form-group
              id="lastName-group"
              label=":اللقب"
              label-for="lastName">

              <b-form-input
                id="lastName"
                v-model="arLastName"
                type="text"
                placeholder="اللقب"
                required
              ></b-form-input>
            </b-form-group>

            <!-- end Last name  -->

            <!-- Arabic Maiden name  -->
            <b-form-group
              id="maidenName-group"
              label=":إسم العائلة قبل الزواج"
              label-for="maidenName">

              <b-form-input
                direction:RTL
                id="maidenName"
                v-model="arMaidenName"
                type="text"
                placeholder="إسم العائلة قبل الزواج"
                required
              ></b-form-input>
            </b-form-group>

            <!-- end Maiden name  -->

          </div>
      </div>

      <div class="row">
        <div class="col">
          <!-- Birthday -->
          <b-form-group
            id="birthday-group"
            label="Date de Naissance:"
            label-for="birthday">

            <vue-bootstrap-datetimepicker
              v-model="birthday"
              :config="options"
            >
            </vue-bootstrap-datetimepicker>
          </b-form-group>
          <!-- end Birthday -->
        </div>
        <div class="col">
          <!-- Birth Country  -->
          <b-form-group
            id="country-group"
            label="Pays De Naissance:"
            label-for="country">

            <b-form-input
              id="country"
              v-model="birthCountry"
              type="text"
              placeholder="Algerie"
              required
            ></b-form-input>
          </b-form-group>
          <!-- end Birth Country  -->
        </div>
        <div class="col">

                <!-- Gender -->
                <b-form-group
                  id="gender-group"
                  label="Sexe:"
                  label-for="gender">

                  <b-form-select
                    class="mb-2 mr-sm-2 mb-sm-0"
                    v-model="gender"
                    :options="[
                      { 'value': 'homme', 'text': 'Homme'},
                      { 'value': 'femme', 'text': 'Femme'},
                    ]"
                    id="gender"
                  >
                  <template slot="first">
                    <option :value="null" disabled>Choisire ...</option>
                  </template>
                  </b-form-select>
                </b-form-group>
                <!-- end Gender -->
        </div>
      </div>

      <div class="row">
        <div class="col">
          <!-- Registred Date -->
          <b-form-group
            id="registredDate-group"
            label="Date d'Inscription:"
            label-for="registredDate">

            <vue-bootstrap-datetimepicker
              v-model="registredDate"
              :config="options"
            >
            </vue-bootstrap-datetimepicker>
          </b-form-group>
          <!-- end Registred Date -->
        </div>

        <div class="col">
          <!-- Phone Number  -->
          <b-form-group
            id="phoneNumber-group"
            label="Numero De Telephone:"
            label-for="phoneNumber">

            <b-form-input
              id="phoneNumber"
              v-model="phoneNumber"
              type="text"
              placeholder="01-23-45-67-89"
              required
            ></b-form-input>
          </b-form-group>
          <!-- end Phone Number  -->
        </div>
      </div>

      <!-- Address  -->
      <b-form-group
        id="address-group"
        label="Adresse:"
        label-for="address">

        <b-form-input
          id="address"
          v-model="addressStreet"
          type="text"
          placeholder="adresse"
          required
        ></b-form-input>
      </b-form-group>
      <!-- end Address  -->

      <div class="row">
        <div class="col">
          <!-- Department  -->
          <b-form-group
            id="department-group"
            label="Daira:"
            label-for="department">

            <b-form-input
              id="department"
              v-model="department"
              type="text"
              placeholder="Bir el djir"
              required
            ></b-form-input>
          </b-form-group>
          <!-- end Department  -->
        </div>
        <div class="col">
          <!-- City  -->
          <b-form-group
            id="city-group"
            label="Wilaya:"
            label-for="city">

            <b-form-input
              id="city"
              v-model="city"
              type="text"
              placeholder="Oran"
              required
            ></b-form-input>
          </b-form-group>
          <!-- end City -->
        </div>
      </div>

      <div class="row">
        <div class="col">
          <!-- File Number -->
          <b-form-group
            id="fileNumber-group"
            label="Numero De Dossier:"
            label-for="fileNumber">

            <b-form-input
              id="fileNumber"
              v-model="fileNumber"
              type="text"
              placeholder="1234"
              required
            ></b-form-input>
          </b-form-group>
          <!-- end File Number -->
        </div>
        <div class="col">
          <!-- job -->
          <b-form-group
            id="job-group"
            label="Travaille:"
            label-for="job">

            <b-form-input
              id="job"
              v-model="job"
              type="text"
              placeholder=""
              required
          ></b-form-input>
          <!-- end Job  -->
        </b-form-group>
        </div>
      </div>
      <b-button
        block type="submit"
        variant="primary">
        Ajouter L'Étudiant
      </b-button>
    </b-form>
  </div>
</template>

<script>
import VueBootstrapDatetimepicker from 'vue-bootstrap-datetimepicker';
import Header from './parts/Header';
import moment from 'moment';

export default {
  name: "addStudent",
  components: {
    VueBootstrapDatetimepicker,
    Header,
  },
  created() {
    //do something after creating vue instance
    this.resetMessagesAlert();
  },
  data: () => ({
    options:{
      format:"YYYY-MM-DD",
      useCurrent: false
    },
    alertMessage:"",
    alertVariant:"",

    student:{},

    firstName: "",
    arFirstName: "",
    lastName: "",
    arLastName: "",
    maidenName: "",
    arMaidenName: "",

    birthday: "",
    birthCountry: "",
    gender: "",

    phoneNumber: "",
    fileNumber: "",
    job: "",
    country: "",
    city: "",
    department: "",
    addressStreet: "",
    registredDate: "",

    studentCreated: false,
    errorCreateStudent: null,
  }),
  methods:{
    // Rset Alert Message
    resetMessagesAlert:function(){
      this.alertVariant = "success";
      this.alertMessage = "L'étudiant a été crée";
    },

    addStudent:function(){
      if(this.errorCreateStudent != null){
        this.resetMessagesAlert();
      }
      if(this.maidenName == ""){
        this.maidenName = "None";
        this.arMaidenName = "None";
      }
      this.student = {
        "first_name": JSON.stringify({"fr":this.firstName, "ar":this.arFirstName}),
        "last_name": JSON.stringify({"fr":this.lastName, "ar":this.arLastName}),
        "maiden_name": JSON.stringify({"fr":this.maidenName, "ar":this.arMaidenName}),
        "phone_number": this.phoneNumber,
        "birthday": moment(this.birthday).format(),
        "birth_city": this.birthCountry,
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
          this.alertMessage = "Il y'a une erreur, l'étudiant n'a pas été crée";
          this.alertVariant = "warning";
        }
        this.studentCreated = true;
    },
    closeAlert:function(){
      this.studentCreated = false;
      this.errorCreateStudent = null;
      this.resetMessagesAlert();
    },
  }
}
</script>
<style scoped>
.ar{
  direction: rtl;
}
</style>