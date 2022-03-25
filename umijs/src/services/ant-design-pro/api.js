// @ts-ignore

/* eslint-disable */
import { request } from 'umi';
/** 获取当前的用户 GET /api/currentUser */

export async function currentUser(options) {
  return request('/api/currentUser', {
    method: 'GET',
    ...(options || {}),
  });
}

export async function outLogin(options) {
  return request('/api/login/outLogin', {
    method: 'POST',
    ...(options || {}),
  });
}

export async function register(payload) {
  return request(`http://localhost:8001/user`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(payload),
    skipErrorHandler: true,
  });
}

export async function login(body) {
  return request('http://localhost:8001/login', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
  });
}
/** 此处后端没有提供注释 GET /api/notices */

export async function getNotices(options) {
  return request('/api/notices', {
    method: 'GET',
    ...(options || {}),
  });
}
/** 获取规则列表 GET /api/rule */

export async function user(params) {
  const options = {
    headers: {
      Authorization: 'Bearer ' + `${localStorage.getItem('token')} `,
    },
  };

  return request('http://localhost:8001/users', {
    method: 'GET',
    params: { ...params },
    ...(options || {}),
  });
}
/** 新建规则 PUT /api/rule */
export async function getRole() {
  return request(`http://localhost:8001/roles`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
      Authorization: 'Bearer ' + `${localStorage.getItem('token')} `,
    },
  });
}

export async function checkProduct(id, payload) {
  console.log(payload);
  console.log(localStorage.getItem('token'));
  return request(`http://localhost:8001/products/${id}/checked`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
      Authorization: 'Bearer ' + `${localStorage.getItem('token')}`,
    },
    body: JSON.stringify(payload),
    skipErrorHandler: true,
  });
}

export async function publishProduct(id, payload) {
  return request(`http://localhost:8001/products/${id}/published`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
      Authorization: 'Bearer ' + `${localStorage.getItem('token')} `,
    },
    body: JSON.stringify(payload),
    skipErrorHandler: true,
  });
}

export async function userDetail(id) {
  console.log(id);
  return request(`http://localhost:8001/user/${id}`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
      Authorization: 'Bearer ' + `${localStorage.getItem('token')} `,
    },
  });
}

export async function updateUser(id, payload) {
  console.log(localStorage.getItem('token'));
  return request(`http://localhost:8001/user/${id}`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
      Authorization: 'Bearer ' + `${localStorage.getItem('token')} `,
    },
    body: JSON.stringify(payload),
    skipErrorHandler: true,
  });
}

export async function removeUser(id) {
  console.log(id);
  return request(`http://localhost:8001/user/${id}`, {
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json',
      Authorization: 'Bearer ' + `${localStorage.getItem('token')} `,
    },
    skipErrorHandler: true,
  });
}

export async function product(params) {
  const options = {
    headers: {
      Authorization: 'Bearer ' + `${localStorage.getItem('token')} `,
    },
  };

  return request('http://localhost:8001/products', {
    method: 'GET',
    params: { ...params },
    ...(options || {}),
  });
}

export async function productDetail(id) {
  return request(`http://localhost:8001/products/${id}`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
      Authorization: 'Bearer ' + `${localStorage.getItem('token')} `,
    },
  });
}

export async function createNewProduct(payload) {
  console.log(payload);
  return request(`http://localhost:8001/product`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      Authorization: 'Bearer ' + `${localStorage.getItem('token')} `,
    },
    body: JSON.stringify(payload),
    skipErrorHandler: true,
  });
}

export async function updateProduct(id, payload) {
  return request(`http://localhost:8001/products/${id}`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
      Authorization: 'Bearer ' + `${localStorage.getItem('token')} `,
    },
    body: JSON.stringify(payload),
    skipErrorHandler: true,
  });
}

export async function removeProduct(id) {
  return request(`http://localhost:8001/product/${id}`, {
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json',
      Authorization: 'Bearer ' + `${localStorage.getItem('token')} `,
    },
    skipErrorHandler: true,
  });
}

export async function updateRule(options) {
  return request('/api/rule', {
    method: 'PUT',
    ...(options || {}),
  });
}
/** 新建规则 POST /api/rule */

export async function addRule(options) {
  return request('/api/rule', {
    method: 'POST',
    ...(options || {}),
  });
}
/** 删除规则 DELETE /api/rule */

export async function removeRule(options) {
  return request('/api/rule', {
    method: 'DELETE',
    ...(options || {}),
  });
}
