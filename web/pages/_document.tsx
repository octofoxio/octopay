import * as React from 'react'
import Document, {Head, Main, NextScript} from 'next/document'
import { extractStyles } from 'evergreen-ui'

export default class MyDocument extends Document {
    props: any
    static async getInitialProps(ctx) {
        const initialProps = await Document.getInitialProps(ctx)
        const { renderPage } = ctx
        const page = renderPage()
        const { css, hydrationScript } = extractStyles()
        return {
            ...page,
            ...initialProps,
            css,
            hydrationScript,


        }
    }

    render() {
        const { css, hydrationScript } = this.props
        return (
            <html>
            <Head>
                <link
                    rel="stylesheet"
                    href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/css/bootstrap.min.css"
                    integrity="sha384-MCw98/SFnGE8fJT3GXwEOngsV7Zt27NXFoaoApmYm81iuXoPkFOJwJ8ERdknLPMO"
                    crossOrigin="anonymous"
                />
                <style dangerouslySetInnerHTML={{ __html: css }} />
                <link
                    rel="stylesheet"
                    href="./static/style.css"
                    />
            </Head>
            <body className="root default-background-color">
            <Main />
            {hydrationScript}
            <NextScript/>
            </body>
            </html>
        )
    }
}