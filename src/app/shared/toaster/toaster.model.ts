import { PositionX, PositionY, ToastPositionModel } from "@syncfusion/ej2-angular-notifications";

/**
 * Defines the type of toaster message
 */
export declare type ToasterType = "success" | "warning" | "info" | "error";

/**
 * Defines the horizontal position of toaster message
 */
export declare type ToasterPositionX = PositionX | number | string;

/**
 * Defines the vertical position of toaster message
 */
export declare type ToasterPositionY = PositionY | number | string;

/**
 * Defines the structure of toaster
 */
export interface ToasterModel {

    /**
     * Type of toaster
     */
    type: ToasterType,

    /**
     * Message / Description to be displayed in the toaster
     */
    message: string,

    /**
     * Title to be displayed in the toaster
     */
    title?: string,

    /**
     * Position to display the toaster
     */
    position?: ToastPositionModel
}
