module.exports = {
    publicPath: '/ui',
    devServer: {
        port: 3030,
        open: true,
        proxy: {
            '/api': {
                target: 'http://127.0.0.1:3000'
            }
        }
    }
}