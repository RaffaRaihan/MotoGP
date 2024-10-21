<template>
    <!-- Start Footer -->
    <footer class="footer">
        <section class="line-1">
            <div class="logo-footer">
                <img src="../assets/img/logo-footer.svg" alt="logo-footer">
            </div>
            <hr>
            <div class="contact">
                <h1>Contact Me</h1>
                <p><font-awesome-icon :icon="['fas', 'envelope']" />raffamraihan580@gmail.com</p>
                <p><font-awesome-icon :icon="['fas', 'phone']" />+62 812-2159-4441</p>
                <p><font-awesome-icon :icon="['fas', 'house']" />Jln. Barupadang Desa Cipada Kec. Cisarua Kab. Bandung Barat Provinsi Jawa Barat</p>
            </div>
            <hr>
            <div class="about-us">
                <h1>About Us</h1>
                <div class="link">
                    <div class="menu-link">
                        <NuxtLink to="/">Home</NuxtLink>
                        <NuxtLink to="/about">About</NuxtLink>
                        <NuxtLink to="/riders">Riders</NuxtLink>
                        <NuxtLink to="/teams">Teams</NuxtLink>
                    </div>
                    <a href="#" class="btn-sign-up-footer">Sign Up</a>
                </div>
            </div>
        </section>
        <section class="line-2">
            <div class="follow">
                <h1>Follow Me</h1>
                <a href="#"><font-awesome-icon :icon="['fab', 'instagram']" />@raffa.mr_</a>
                <a href="#"><font-awesome-icon :icon="['fab', 'facebook']" />Raffa Raihan</a>
                <a href="#"><font-awesome-icon :icon="['fab', 'x-twitter']" />Raffa Raihan</a>
                <a href="#"><font-awesome-icon :icon="['fab', 'linkedin']" />Raffa Raihan</a>
                <a href="#"><font-awesome-icon :icon="['fab', 'github']" />RaffaRaihan</a>
            </div>
            <p class="copyright">&copy;raffamr2024</p>
        </section>
    </footer>
    <!-- End Footer -->
</template>

<style scoped>
/* footer start */
.footer{
    margin-top: 80px;
    padding-bottom: 15px;
    background-color: #0B192C;
    color: white;
}

/* line-1 */
.line-1{
    display: flex;
    margin-inline: 100px;
    padding-top: 22px;
    height: 168px;
    justify-content: space-between;
}
.logo-footer{
    width: 212px;
    height: 168px;
}

.logo-footer img{
    width: 212px;
    height: 132px;
    padding-top: 18px;
    padding-bottom: 18px;
}

.contact{
    width: 369px;
    height: 168px;
    align-items: start;
}

.contact h1{
    font-family: 'MotoGP Display';
    font-weight: bold;
    font-size: 40px;
}

.contact p{
    font-family: 'MotoGP Text';
    font-weight: 300;
    font-style: normal;
    margin-top: 12px;
}
.contact p i{
    margin-right: 10px;
}

.about-us{
    width: 338px;
    height: 168px;
    align-items: start;
}

.about-us h1{
    font-family: 'MotoGP Display';
    font-weight: bold;
    font-size: 40px;
}

.link {
    display: flex;
    justify-content: space-between;
}

.menu-link a{
    display: block;
    font-family: 'MotoGP Text';
    font-weight: 300;
    font-style: normal;
    text-decoration: none;
    color: white;
    margin-top: 10px;
}
 .menu-link a:hover{
    width: 33px;
    font-family: 'MotoGP Text';
    font-weight: 300;
    font-style: italic;
    color: #CD1818;
    text-decoration: underline;
}

.btn-sign-up-footer{
    border-radius: 15px;
    font-family: 'MotoGP Text' ,sans-serif;
    font-weight: normal;
    font-size: 16px;
    background-color: #CD1818;
    padding: 11px 25px 11px 25px ;
    border: solid;
    text-decoration: none;
    color: #000;
    margin-top: 85px;
}

/* line-2 */

.line-2{
    display: flex;
    margin-inline: 100px;
    margin-top: 44px;
    margin-bottom: 15px;
    justify-content: space-between;
}

.follow{
    height: 80px;
}

.follow h1{
    font-family: 'MotoGP Display';
    font-weight: bold;
    font-size: 40px;
}
.follow a{
    font-family: 'MotoGP Text';
    font-weight: 300;
    font-style: normal;
    text-align: center;
    text-decoration: none;
    color: white;
    margin-right: 12px;
}
.follow a i {
    margin-right: 10px;
}

.follow a:hover{
    font-family: 'MotoGP Text';
    font-weight: 300;
    font-style: italic;
    text-decoration: underline;
    color: #CD1818;
}

.copyright{
    font-family: 'MotoGP Text';
    font-weight: 300;
    font-style: normal;
    margin-top: 55px;
}
/* footer end */
</style>