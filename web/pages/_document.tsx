import Document, {Head, Main, NextScript} from 'next/document'

export default class MyDocument extends Document {
    static async getInitialProps(ctx) {
        const initialProps = await Document.getInitialProps(ctx)
        return {...initialProps}
    }

    render() {
        return (
            <html>
            <Head>
                <link
                    rel="stylesheet"
                    href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/css/bootstrap.min.css"
                    integrity="sha384-MCw98/SFnGE8fJT3GXwEOngsV7Zt27NXFoaoApmYm81iuXoPkFOJwJ8ERdknLPMO"
                    crossOrigin="anonymous"
                />
                <link
                    rel="stylesheet"
                    href="./static/style.css"
                    />
            </Head>
            <body className="root default-background-color">
            <Main />
            <NextScript/>
            </body>
            </html>
        )
    }
}