export default [
  {
    path: '/user',
    layout: false,
    routes: [
      {
        path: '/user',
        routes: [
          {
            name: 'login',
            path: '/user/login',
            component: './user/Login',
          },
          {
            name: 'register',
            path: '/user/register',
            component: './user/Register',
          },
        ],
      },
      {
        component: './404',
      },
    ],
  },
  {
    path: '/welcome',
    name: 'welcome',
    icon: 'smile',
    component: './Welcome',
  },
  {
    name: 'list.table-list',
    icon: 'table',
    path: '/admin/users',
    access: 'canAdmin',
    component: './admin/UserTable',
  },
  {
    icon: 'table',
    name: 'list.basic-list',
    path: '/admin/products',
    access: 'canAdmin',
    component: './admin/ProductTable',
  },

  {
    name: 'list.table-list',
    icon: 'table',
    path: '/users',
    routes: [
      { access: 'canViewer' },
      { access: 'canSigner' },
      { access: 'canMaker' },
      { access: 'canChecker' },
    ],

    component: './UserTable',
  },
  {
    icon: 'table',
    name: 'list.basic-list',
    path: '/productlist',
    access: 'canViewer',
    component: './viewer/ProductTable',
  },
  {
    icon: 'table',
    name: 'list.basic-list',
    path: '/checker/productlist',
    access: 'canChecker',
    component: './checker/ProductTable',
  },
  {
    icon: 'table',
    name: 'list.basic-list',
    path: '/signer/productlist',
    access: 'canSigner',
    component: './signer/ProductTable',
  },
  {
    path: '/',
    redirect: '/welcome',
  },
  {
    component: './404',
  },
];
