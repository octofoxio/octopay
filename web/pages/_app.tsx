import * as React from 'react'

import {Pane} from 'evergreen-ui'
import {ApolloProvider} from 'react-apollo'
import {apolloClient, ApolloClientContext} from "../utils/apollo";
import App, {Container} from 'next/app'

const NavigationBar = () => {
    return (
        <Pane
            height={79}
            borderBottom
            backgroundColor="var(--white)"
            className={"navigation-bar"}
        >
        </Pane>
    )
}
export default class MyApp extends App {
    static async getInitialProps({Component, ctx}) {
        let pageProps = {}

        if (Component.getInitialProps) {
            pageProps = await Component.getInitialProps(ctx)
        }

        return {pageProps}
    }

    render() {
        const {Component, pageProps} = this.props
        return (
            <Container>

                <ApolloClientContext.Provider value={apolloClient}>
                    <ApolloProvider client={apolloClient}>

                        <Pane height={"100vh"} display="flex" flexDirection="column">
                            <NavigationBar/>
                            <Component {...pageProps} />
                        </Pane>
                    </ApolloProvider>
                </ApolloClientContext.Provider>
            </Container>
        )
    }
}