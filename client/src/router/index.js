import Vue from "vue";
import Router from "vue-router";

Vue.use(Router);

export default new Router({
    routes: [
        {
            path: "/",
            component: () => import("@/views/Home")
        // },
        // {
        //     name: "games",
        //     path: "/games/all",
        //     component: () => import("@/views/Games"),
        //     props: true
        // },
        // {
        //     name: "heroes",
        //     path: "/heroes/:game",
        //     component: () => import("@/views/Heroes"),
        //     props: true
        }
    ]
});