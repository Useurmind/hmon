import { Button, createStyles, Theme, WithStyles } from "@material-ui/core";
import withStyles from "@material-ui/core/styles/withStyles";
import * as React from "react";
import { useCallback } from "react";
import { IUseStoreFromContainerContextProps, useStoreStateFromContainerContext } from "rfluxx-react";
import { IRouterStore, IRouterStoreState } from "rfluxx-routing";

export const styles = (theme: Theme) => createStyles({
    root: {
    }
});

/**
 * Props for { @RouteButton }
 */
export interface IRouteButtonProps
    extends WithStyles<typeof styles>
{
    /**
     * Link to which to route.
     */
    href: string;
}

/**
 * Implementation of RouteButton.
 */
const RouteButtonImpl: React.FunctionComponent<IRouteButtonProps> = props => {
    const { classes } = props;

    const [ storeState, store ] = useStoreStateFromContainerContext<IRouterStore, IRouterStoreState>({
        storeRegistrationKey: "IRouterStore"
    });
    const routeToLink = useCallback(() => store.navigateToPath.trigger(props.href), [ store ]);

    if (!storeState)
    {
        return null;
    }

    return <Button onClick={routeToLink}>{props.children}</Button>;
};

/**
 * Component that shows a button which routes to a link on the page.
 */
export const RouteButton = withStyles(styles)(RouteButtonImpl);
