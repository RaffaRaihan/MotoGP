<template>
    <div>
      <Navigasi />
      <main class="main-about">
        <div class="title-about">
          <h1>Ini adalah halaman about tentang tugas akhir saya</h1>
        </div>
      </main>
      <Footer />
    </div>
</template>
  
<style scoped>
.main-about {
  background: url('@/assets/img/main-about.png');
  height: 810px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.title-about h1 {
  font-family: 'MotoGP Display', sans-serif;
  font-size: 40px;
  color: white;
  text-align: center;
  max-width: 800px;
  margin: 0 auto;
  padding: 20px;
  background-color: rgba(0, 0, 0, 0.7);
  border-radius: 10px;
}

@media (max-width: 768px) {
  .main-about {
    height: 500px;
  }

  .title-about h1 {
    font-size: 28px;
  }
}
</style>