import * as React from 'react'
import { Dialog } from 'evergreen-ui'
import {UseUpdatePaymentStatus} from "../hooks/update-tx-status";


export const ConfirmPaymentDialog = (title: string, onConfirm: () => void) => {
    const [isShown, setIsShown]= React.useState(false)
    const { isLoading, mutate } = UseUpdatePaymentStatus()
    function Done() {
        setIsShown(false)
    }
    return (
        <Dialog
            isShown={isShown}
            title="Loading confirmation"
            onCloseComplete={Done}
            onConfirm={() => mutate}
            isConfirmLoading={isLoading}
            confirmLabel={isLoading ? 'Loading...' : 'Confirm Loading'}
        >
            Dialog content
        </Dialog>
    )
}