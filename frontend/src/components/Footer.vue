<template>
    <div class="footer">
        <v-footer dark lighten-1 padless>
            <v-col
                class="dark lighten-3 text-center"
                cols="12"
            >
                <v-btn
                    v-for="(social, index) in socials"
                    :key="index"
                    class="mx-4 white--text justify-center text-center"
                    icon
                >
                    <a :href="social.link" target="_blank">
                        <v-icon size="24px">
                        {{ social.icon }}
                        </v-icon>
                    </a>
                </v-btn>
            </v-col>
            <v-col
                class="dark lighten-3 text-center"
                cols="12"
            >
                <strong>Yource</strong> - Community like slack.<br>
                © {{ new Date().getFullYear() }} -  Diwamoto.
            </v-col>
        </v-footer>
    </div>
</template>

<script>
export default {
    data: () => ({
    socials: [
        { icon: 'mdi-facebook', link: 'https://www.facebook.com/profile.php?id=100045347584995', },
        { icon: 'mdi-twitter', link: 'https://twitter.com/Diwamoto_', },
        { icon: 'mdi-github', link: 'https://github.com/Diwamoto', },
    ],
  }),
}
</script>

<style>
.footer a{
    color: white;
    text-decoration: none;
}
</style>