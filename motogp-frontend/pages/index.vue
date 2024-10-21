<template>
  <div>
    <Navigasi />
    <section class="main">
      <div class="container">
        <h1>Hallo, <span>user</span></h1>
        <p>Selamat datang di Website list Rider Motogp, 
          Kalian bisa melihat para Rider Motogp dari Moto 3, Moto 2, dan Motogp.
        </p>
        <NuxtLink to="/riders" class="get-started">
          Mulai Sekarang <font-awesome-icon :icon="['fas', 'arrow-right']" />
        </NuxtLink>
      </div>
      <div class="hero">
        <img src="@/assets/img/Pecco.png" alt="hero">
      </div>
    </section>
    <Footer />
  </div>
</template>

<style scoped>
/* main index start */
.main{
    display: flex;
    margin-inline: 100px;
    height: 700px;
    margin-top: 30px;
    justify-content: space-between;
}

.container{
    width: 541px;
    height: 444px;
    margin-top: 70px;
}

.container h1{
    font-family: 'MotoGP Text';
    font-weight: bold;
    font-style: normal;
    font-size: 40px;
    margin-top: 15px;
    margin-bottom: 40px;
}

.container h1 span{
    color: #CD1818;
}

.container p{
    font-family: 'MotoGP Text';
    font-weight: 300;
    font-style: normal;
    font-size: 32px;
    margin-bottom: 30px;
}

.get-started{
    padding: 4px 10px 4px 10px;
    text-decoration: none;
    background-color: #CD1818;
    color: #000;
    border-radius: 12px;
    font-family: 'MotoGP Display';
    font-weight: bold;
    font-style: normal;
    font-size: 40px;
    text-align: start;
}

.hero{
    width: 620px;
    height: 700px;
}
/* main index end */
</style>