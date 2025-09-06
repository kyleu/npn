import "./client.css";
import { appInit } from "./app";
import { audit } from "./audit";
import { autocompleteInit } from "./autocomplete";
import { flashInit } from "./flash";
import { formInit } from "./form";
import { linkInit } from "./link";
import { menuInit } from "./menu";
import { modalInit } from "./modal";
import { modeInit } from "./mode";
import { SocketMessage, socketInit } from "./socket";
import { tagsInit } from "./tags";
import { themeInit } from "./theme";
import { timeInit } from "./time";

declare global {
  // eslint-disable-next-line @typescript-eslint/consistent-type-definitions
  interface Window {
    npn: {
      wireTime: (el: HTMLElement) => void;
      relativeTime: (el: HTMLElement) => string;
      autocomplete: (
        el: HTMLInputElement,
        url: string,
        field: string,
        title: (x: unknown) => string,
        val: (x: unknown) => string
      ) => void;
      setSiblingToNull: (el: HTMLElement) => void;
      initForm: (frm: HTMLFormElement) => void;
      flash: (key: string, level: "success" | "error", msg: string) => void;
      tags: (el: HTMLElement) => void;
      Socket: unknown;
    };
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    audit: (s: string, ...args: any) => void;
  }
}

export function init(): void {
  const [s, i] = formInit();
  const [wireTime, relativeTime] = timeInit();
  window.npn = {
    wireTime: wireTime,
    relativeTime: relativeTime,
    autocomplete: autocompleteInit(),
    setSiblingToNull: s,
    initForm: i,
    flash: flashInit(),
    tags: tagsInit(),
    Socket: socketInit()
  };
  menuInit();
  modeInit();
  linkInit();
  modalInit();
  themeInit();
  window.audit = audit;
  appInit();
}

document.addEventListener("DOMContentLoaded", init);
