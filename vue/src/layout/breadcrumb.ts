import {ref} from "vue";

export interface Breadcrumb {
  readonly title: string;
  readonly path: string;
}

export const breadcrumbsRef = ref<Breadcrumb[]>();
