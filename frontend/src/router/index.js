import Vue from 'vue'
import Router from 'vue-router'
import ListStudents from '@/components/ListStudents.vue'
import AddStudents from '@/components/AddStudent.vue'
import AddExamList from '@/components/AddExamList.vue'
import ExamLists from '@/components/ExamLists.vue'
import StudentDetails from '@/components/StudentDetails.vue'

import Datetime from 'vue-datetime'
import 'vue-datetime/dist/vue-datetime.css'

Vue.use(Router)
Vue.use(Datetime)
export default new Router({
  routes:[
    {path: "/", name: "listStudents",component: ListStudents},
    {path: "/add-student", name: "addStudent",component: AddStudents},
    {path: "/add-exam-list", name: "addExamList",component: AddExamList},
    {path: "/exam-lists", name: "examLists",component: ExamLists},
    {path: "/:id", name: "studentDetails",component: StudentDetails}
  ]
})