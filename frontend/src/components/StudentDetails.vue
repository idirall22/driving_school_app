<template>
  <div v-if="ready" id="studentDetails">
    <b-alert
      v-model="studentUpdated"
      :variant="alertVariant"
      dismissible
      @dismissed="closeAlert()">
      {{alertMessage}}
    </b-alert>
      <Header initTitle="Détails"></Header>

      <p>Prochain Examen:</p>
      <b-progress :max="3" class="mb-3" height= "2rem">
        <b-progress-bar
          variant="success"
          :value="getStudentNextExamPourcent()"
        > <strong>{{getStudentNextExamName()}}</strong>
        </b-progress-bar>
      </b-progress>

    <b-form @submit="editStudent">
      <div class="row">
          <div class="col">
            <!-- First name  -->
            <b-form-group
              id="firstName-group"
              label="Prénom:"
              label-for="firstName">

            <b-form-input
              id="firstName"
              v-model="student.first_name"
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
                v-model="student.last_name"
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
                v-model="student.maiden_name"
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
                v-model="student.ar_first_name"
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
                v-model="student.ar_last_name"
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
                v-model="student.ar_maiden_name"
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
              v-model="student.birth_city"
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
                    v-model="student.gender"
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
              v-model="student.phone_number"
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
          v-model="student.address_street"
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
              v-model="student.department"
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
              v-model="student.city"
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
              v-model="student.file_number"
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
              v-model="student.job"
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
        Modifier L'Étudiant
      </b-button>
    </b-form>

  </div>
</template>


<script>
import VueBootstrapDatetimepicker from 'vue-bootstrap-datetimepicker';
import Header from './parts/Header';
import moment from 'moment';

export default {
  name: "studentDetails",
  components: {
    VueBootstrapDatetimepicker,
    Header
  },
  mounted: function(){
    this.getStudent();
    this.resetMessagesAlert();
  },
  data: () => ({
    ready: false,
    options:{
      format:"DD-MM-YYYY",
      useCurrent: false
    },
    data: {},
    registredDate: null,
    birthday: null,
    studentUpdated: false,
    alertVariant: "",
    alertMessage: "",
    errorGetStudent: null,
    errorUpdate: null,
    student:{},
    errrrr: null,
    id: null,
  }),
  methods:{
    closeAlert: function(){
      this.studentUpdated = false;
      this.errorGetStudent = null;
    },
    resetMessagesAlert:function(){
      this.alertVariant = "success";
      this.alertMessage = "L'étudiant a été modifié";
    },
    getDataMultiLanguage:function(){
      try {
        let fn = JSON.parse(this.student.first_name);
        this.student.first_name = fn.fr;
        this.student.ar_first_name = fn.ar;
      } catch (e) {
        this.errrrr = e;
        this.student.first_name = "";
        this.student.ar_first_name = "";
      }
      try {
        let ln = JSON.parse(this.student.last_name);
        this.student.last_name = ln.fr;
        this.student.ar_last_name = ln.ar;
      } catch (e) {
        this.errrrr = e;

        this.student.last_name = "";
        this.student.ar_last_name = "";
      }
      try {
        let mn = JSON.parse(this.student.maiden_name);
        this.student.maiden_name = mn.fr;
        this.student.ar_maiden_name = mn.ar;
      } catch (e) {
        this.errrrr = e;

        this.student.maiden_name = "";
        this.student.ar_maiden_name = "";
      }
    },
    getStudent: function(){
      this.id = this.$route.params.id;
      window.backend.Service.GetStudent(this.id, "").then(
        data=>{
          this.student = data;
          this.getDataMultiLanguage();
          this.registredDate = moment(this.student.registred_date).format();
          this.birthday = moment(this.student.birthday).format();
          this.ready = true;
        },
        err=>{
          this.errorGetStudent = err;
          this.alertVariant = "warning";
          this.alertMessage = "Erreur du Server";
          this.studentUpdated = true;
        }
      );
    },
    editStudent: function(){
      this.student.birthday = moment(this.birthday).format();
      this.student.registred_date = moment(this.registredDate).format();
      try {
        this.student.first_name = JSON.stringify(
          {"fr": this.student.first_name, "ar": this.student.ar_first_name}
        );

        this.student.last_name = JSON.stringify(
          {"fr": this.student.last_name, "ar": this.student.ar_last_name}
        );

        this.student.maiden_name = JSON.stringify(
          {"fr": this.student.maiden_name, "ar": this.student.ar_maiden_name}
        );

        // delete this.student["ar_first_name"];
        // delete this.student["ar_last_name"];
        // delete this.student["ar_maiden_name"];

      } catch (e) {
        this.studentUpdated = true;
        this.alertVariant = "warning";
        this.alertMessage = "Erreur L'étudiant n'a pas été modifié";
      }
      window.backend.Service.UpdateStudent(this.student).then(
        data=>{
          this.data = data;
        },
        err =>{
          this.errorUpdate = err;
        }
      );if(this.errorUpdate != null){
        this.alertVariant = "warning";
        this.alertMessage = "Erreur L'étudiant n'a pas été modifié";
      }
      this.studentUpdated = true;
      this.getDataMultiLanguage();
    },
    getStudentNextExamName: function(){
      switch (this.student.next_exam) {
        case "Highway code":
          return "Code";
        case "Niche":
          return "Créneau";
        case "Circuit":
          return "Circuit";
      }
    },
    getStudentNextExamPourcent: function(){
      switch (this.student.next_exam) {
        case "Highway code":
          return 1;
        case "Niche":
          return 2
        case "Circuit":
          return 3
      }
    }
  }
}
</script>

<style scoped>
</style>
