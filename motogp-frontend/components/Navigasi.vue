<template>
    <nav class="navigasi">
      <div class="logo">
        <img src="@/assets/img/logo.svg" alt="logo">
        <h1>MotoGP</h1>
      </div>
      <div class="nav-items">
        <div class="nav-link">
          <NuxtLink to="/">Home</NuxtLink>
          <NuxtLink to="/about">About</NuxtLink>
          <NuxtLink to="/riders">Riders</NuxtLink>
          <NuxtLink to="/teams">Teams</NuxtLink>
        </div>
        <a href="#" class="btn-sign-up">Sign Up</a>
      </div>
    </nav>
  </template>
  
  <style scoped>
  /* navigasi start */
.navigasi{
    display: flex;
    width: auto;
    height: 100px;
    margin-inline: 100px;
    margin-top: 20px;
    justify-content: space-between;
}

.logo{
    display: flex;
    width: 270px;
    height: 100px;
    font-family: 'MotoGP Display', sans-serif;
    align-items: center;
    gap: 25px;
}

.logo img{
    width: 120px;
    height: 64px;
    padding: 0px,18px;
}

.logo h1 {
    font-size: 40px;
}

.nav-items{
    display: flex;
    height: 45px;
    font-family: 'MotoGP Text' ,sans-serif;
    margin-top: 28px;
    margin-bottom: 27px;
}

.nav-items .nav-link{
    display: flex;
    gap: 45px;
    align-items: center;
}

.nav-link a{
    text-decoration: none;
    color: #000;
    font-size: 16px;
}

.nav-link a:active{
    color: #CD1818;
}

.nav-link a:hover{
    color: #CD1818;
}

.btn-sign-up{
    margin-left: 45px;
    border-radius: 15px;
    font-family: 'MotoGP Text';
    font-weight: normal;
    font-size: 16px;
    background: #CD1818;
    padding: 11px 25px 11px 25px;
    border: solid;
    text-decoration: none;
    color: #000;
}

.btn-sign-up :hover{
    background: transparent;
    color: #CD1818;
}
/* navigasi end */
  </style>