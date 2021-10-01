import Vue from 'vue'
import Router from 'vue-router'
import HomePage from '@/components/HomePage'
import RequestPage from '@/components/RequestPage'
import SubmissionPage from '@/components/SubmissionPage'
import MyPage from '@/components/MyPage'

Vue.use(Router)

export default new Router({
    mode: 'history',
    routes: [{
            path: '/',
            name: 'HomePage',
            component: HomePage
        },
        {
            path: '/request/:request_id',
            name: 'RequestPage',
            component: RequestPage
        },
        {
            path: '/submission',
            name: 'SubmissionPage',
            component: SubmissionPage
        },
        {
            path: '/user/:user_id',
            name: 'MyPage',
            component: MyPage
        }
    ]
});