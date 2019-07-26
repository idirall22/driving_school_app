<template>
  <div id="studentDetails">

    <div v-if="ready">
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
          :value="student.next_exam"
        > <strong>{{student.getExamName(student.next_exam)}}</strong>
        </b-progress-bar>
      </b-progress>

      <div v-if="student != null">
        <b-form @submit="updateStudent">
          <div class="row">
              <div class="col">
                <!-- First name  -->
                <b-form-group
                  id="firstName-group"
                  label="Prénom:"
                  label-for="firstName">

                <b-form-input
                  id="firstName"
                  v-model="student.first_name_fr"
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
                    v-model="student.last_name_fr"
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
                    v-model="student.maiden_name_fr"
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
                    v-model="student.first_name_ar"
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
                    v-model="student.last_name_ar"
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
                    v-model="student.maiden_name_ar"
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
                  v-model="student.birthday"
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
                  v-model="student.country"
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
                  v-model="student.registred_date"
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
      <br>
      <Header initTitle="Examens Passé"></Header>
      <div class="table-responsive">
        <table class="table table-striped table-sm">
          <thead>
            <tr>
              <th class="align-middle text-center">N°</th>
              <th class="align-middle text-center">Date d'éxamen</th>
              <th class="align-middle text-center">Examen</th>
              <th class="align-middle text-center">Résultats</th>
            </tr>

          </thead>
          <tbody>
            <tr v-for="(exam, index) in student.exams" :key="exam.id">
              <td class="align-middle text-center">{{index+1}}</td>
              <td class="align-middle text-center">{{exam.date_exam | moment2}}</td>
              <td class="align-middle text-center">{{student.getExamName(exam.exam)}}</td>
              <td v-if="exam.status" class="align-middle text-center">Gagné</td>
              <td v-else class="align-middle text-center">Perdu</td>
            </tr>
          </tbody>
        </table>
      </div>

    </div>
  </div>

</template>


<script>
import VueBootstrapDatetimepicker from 'vue-bootstrap-datetimepicker';
import Header from './parts/Header';
import Student from './service/student.js'
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
    options:{
      format:"DD-MM-YYYY",
      useCurrent: false
    },
    ready: false,
    studentUpdated: false,
    alertVariant: "",
    alertMessage: "",
    errorGetStudent: null,
    errorUpdate: null,

    student: {},
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
    getStudent: function(){
      this.id = this.$route.params.id;
      window.backend.Service.GetStudent(this.id, "").then(
        data=>{
          this.student = new Student(data["student"], data["exams"]);
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
    updateStudent: function(){
      this.student.outStudent();

      window.backend.Service.UpdateStudent(
        this.student)
        .then(
          data=>{
            this.data = data;
        },
        err =>{
          this.errorUpdate = err;
        }
      );
      if(this.errorUpdate != null){
        this.alertVariant = "warning";
        this.alertMessage = "Erreur L'étudiant n'a pas été modifié";
      }
      this.studentUpdated = true;
    },
  }
}
</script>

<style scoped>
</style>
