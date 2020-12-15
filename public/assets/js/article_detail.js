const article = new Vue({
  el: '#article-detail',
  data: {
    user: '',
    article: {},
  },
  methods: {
    getArticle() {
      const id = location.href.match(/articles\/(\d+)/)[1];
      const url = `/api/articles/${id}`;
      const headers = {'Authorization': `Bearer ${this.getToken()}`};
      const _this = this;

      fetch(url, {headers}).then(response => {
        if(response.ok) {
          return response.json()
        }
        return {}
      }).then(json => {
        _this.article = json;
      })
    },
    deleteArticle(id) {
      const url = `/api/articles/${id}`;
      const method = 'DELETE';
      const headers = {'Authorization': `Bearer ${this.getToken()}`};

      fetch(url, {method, headers}).then(response => {
        if(response.ok) {
          location.href = '/articles';
        }
      })
    },
    updateArticle(id) {
      const url = `/api/articles/${id}`;
      const method = 'PUT';
      const headers = {
        'Authorization': `Bearer ${this.getToken()}`,
        'Content-Type': 'application/json; charset=UTF-8'
      };
      const body = JSON.stringify({
        title: this.article.title,
        content: this.article.content,
      });

      fetch(url, {method, headers, body})
    },
    getToken() {
      return localStorage.getItem('token');
    },
    logout() {
      localStorage.removeItem('token');
      location.href = '/';
    },
  },
  created() {
    const date = new Date();
    const claims = JSON.parse(atob(this.getToken().split('.')[1]));
    this.user = claims.name;
    if(claims.exp < Math.floor(date.getTime() / 1000)) {
      this.logout();
    } else {
      this.getArticle();
    }
  },
});
