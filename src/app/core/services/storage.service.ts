import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class StorageService {
  private storage = localStorage || window.localStorage
  constructor() { }
  setItem(key: string, value: any) {
    this.storage.setItem(key, value)
  }
  getItem(key: string): any {
    return this.storage.getItem(key)
  }
  hasItem(key: string): boolean {
    return !!this.storage.getItem(key)
  }
  removeItem(key: string) {
    this.storage.removeItem(key)
  }
  clear() {
    this.storage.clear()
  }
}