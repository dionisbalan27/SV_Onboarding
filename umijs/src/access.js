/**
 * @see https://umijs.org/zh-CN/plugins/plugin-access
 * */
export default function access(initialState) {
  const { currentUser } = initialState ?? {};
  console.log(currentUser);
  return {
    canAdmin: currentUser && currentUser.access === 'admin',
    canViewer: currentUser && currentUser.access === 'viewer',
    canChecker: currentUser && currentUser.access === 'checker',
    canMaker: currentUser && currentUser.access === 'maker',
    canSigner: currentUser && currentUser.access === 'signer',
  };
}
