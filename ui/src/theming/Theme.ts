import { createMuiTheme, PaletteType } from "@material-ui/core";

// customized your theme here
declare module "@material-ui/core/styles/createMuiTheme" {
    interface Theme {
    }

    // allow configuration using `createMuiTheme`
    interface ThemeOptions {
    }
}

export const createTheme = (name: string) => {
    // here you could implement creation of different themes depending on the name
    return createMuiTheme();
}