var vm = new Vue({
    el: '#root',
    data: {
        movies: [],
        modal: false,
        loginForm: false,
        signUpForm: false,
        movieForm: false,
        displayUser: false,
        logged: false,
        connectedUserName: '',
        connectedUserID: '',
        currentUserNameMovies: '',
        currentUserID: '',
        modalTitle: '',
        sEmail: '',
        sPassword: '',
        sName: '',
        sTitle: '',
        sDescription: ''
    },
    created: function(){
        this.userStatus();
        this.getMoviesSortedBy("created", "");
    },
    methods: {
        closeModal: function(){
            this.modal = false;
        },
        userStatus: function(){
            this.$http.get("users/session").then(function (response) {
                    var self = this;
                    response.json().then(function(data) {
                        self.connectedUserID = data.user_id;
                        self.connectedUserName = data.user_name;
                        self.logged = data.logged_in;
                    });
            });
        },
        getMoviesSortedBy: function(sorting, user){

            this.$http.get("movies/" + sorting + "/" + user + "/").then(function (response) {
                var self = this;

                self.displayUser = (user.length == 0) ? false : true;
                self.currentUserID = (user.length != 0) ? user : '';
                self.movies = [];

                response.json().then(function(data) {

                    if (response.status === 200){
                        self.movies = data;
                        self.currentUserNameMovies = (user.length != 0) ? data[0].user.name : '';
                        window.scrollTo(0, 0);
                    }

                });

            }).catch(function(response){
                this.catchError(response);
            });

        },
        submitMovie: function(title, description){

            this.$http.post("movies", {title: title, description: description}).then(function(response){
                var self = this;

                response.json().then(function(data){

                    self.displayUser = false;
                    self.movies = [];
                    self.getMoviesSortedBy('created', '');

                    vm.movies.unshift(
                        {
                            title: data.title,
                            description: data.description,
                            created: "just now",
                            user_id: data.user_id,
                            user: {
                                name: data.user.name
                            }
                        }
                    );
                    self.modal = false;
                    self.catchError(response, "Thank you for your movie submission!");
                });

            }).catch(function(response){
                this.catchError(response);
            });
        },
        voteMovie: function(movie_id, action, index){
            var self = this;

            var positive = (action === "like") ? 1 : 0;

            this.$http.post("movies/" + movie_id + "/vote/" + positive).then(function(response){

                response.json().then(function(data){
                    self.movies.splice(index, 1, data);
                    self.catchError(response, "Thank you for voting!");
                })

            }).catch(function(response){
                self.catchError(response);
            });
        },
        retractVote: function(movie_id, index){

            this.$http.delete("movies/" + movie_id + '/vote').then(function(response){

                response.json().then(function(data){
                   vm.movies.splice(index, 1, data);
                });

            }).catch(function(response){
                self.catchError(response);
            });
        },
        signIn: function(email, password){

            this.$http.post("users/session", {email: email, password: password}).then(function(response){

                if (response.status === 200){
                    location.reload();
                }

            }).catch(function(response){
                this.catchError(response);
            });

        },
        signUp: function(email, password, name){

            this.$http.post("users", {email: email, password: password, name: name}).then(function(response){

                if (response.status === 201){
                    location.reload();
                }

            }).catch(function(response){
                this.catchError(response);
            });

        },
        signOut: function(){

            this.$http.delete("/users/session").then(function(response){

                if (response.status === 200){
                    location.reload();
                }

            }).catch(function(response){
                console.log(response);
            });
        },
        showLoginModal: function(){
            this.modal = true;
            this.loginForm = true;
            this.signUpForm = false;
            this.movieForm = false;
            this.modalTitle = "Log In";
        },
        showSignUpModal: function(){
            this.modal = true;
            this.signUpForm = true;
            this.loginForm = false;
            this.movieForm = false;
            this.modalTitle = "Sign Up"
        },
        showMovieModal: function(){
            this.modal = true;
            this.movieForm = true;
            this.loginForm = false;
            this.signUpForm = false;
            this.modalTitle = "Submit Movie"
        },
        catchError: function (response, message) {

            switch (response.status){
                case 400:
                case 404:
                case 401:
                case 409:
                    Vue.toasted.show(response.body.message, {
                        theme: "primary",
                        position: "bottom-left",
                        duration: 2500,
                        type: 'error',
                        icon: 'exclamation-circle'
                    });
                break;
                case 200:
                case 201:
                    Vue.toasted.show(message, {
                        theme: "primary",
                        position: "bottom-left",
                        duration: 2500,
                        type: 'success',
                        icon: 'check'
                    });
                break;
            }

        }
    }
});
