let app = new Vue({
    el: "#app",
    vuetify: new Vuetify(),
    data: {
        anime: null,
        selectValue: "",
        items: [],
        selected: {
        },
        links: [
            'Главная',
            'Ничего',
        ],
        libs: [{
            "user": "Влад",
            "url": "/vlad/",
            "selected": ["Смотрю", "Буду смотреть", "Просмотрено"],
        }, {
            "user": "Иван",
            "url": "/ivan/",
            "selected": ["Смотрю", "watching","Просмотрено", "completed"]
        }]
    },
    computed: {},
    methods: {
        changeUser(i) {
            this.selected = this.libs[i];
        },
        selectCategory(e) {
            let v = this;
            if (v.selected.user == v.libs[0].user) {
                let id;
                v.libs[0].selected.forEach((el, i)=> {
                    if (el == e) {
                        id = i+1;
                    }
                });
                v.selectValue = id;
                v.getAnime();
            } 
        },
        getAnime() {
            let xhr = new XMLHttpRequest();
            xhr.onload = function () {
                var res = JSON.parse(xhr.response);
                anime = res;
            };
            xhr.open('GET', this.selected.url + this.selectValue, true);
            xhr.send();
        },
        refresh() {
            this.getAnime();
        }
    },
    mounted() {
        this.selected = this.libs[0];
        this.items = this.selected.selected;
    }
});