import Vue from 'vue'
import Router from 'vue-router'
import ListStudents from '@/components/ListStudents.vue'
import AddStudents from '@/components/AddStudent.vue'
import AddExamList from '@/components/AddExamList.vue'
import ExamLists from '@/components/ExamLists.vue'
import StudentDetails from '@/components/StudentDetails.vue'

import 'bootstrap/dist/css/bootstrap.css';
import 'pc-bootstrap4-datetimepicker/build/css/bootstrap-datetimepicker.css';
import moment from 'moment';


Vue.use(Router)
Vue.filter("capitalize",str => str.charAt(0).toUpperCase() + str.slice(1))
Vue.filter("moment",date => moment(date).format('D MMMM YYYY'))


export default new Router({
  routes:[
    {path: "/", name: "listStudents",component: ListStudents},
    {path: "/add-student", name: "addStudent",component: AddStudents},
    {path: "/add-exam-list", name: "addExamList",component: AddExamList},
    {path: "/exam-lists", name: "examLists",component: ExamLists},
    {path: "/:id", name: "studentDetails",component: StudentDetails}
  ]
})