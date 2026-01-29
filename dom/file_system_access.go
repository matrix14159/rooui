package dom

// This file implements File System Access API
// https://developer.mozilla.org/en-US/docs/Web/API/File_System_Access_API

import (
	"syscall/js"
)

// FileSystemHandle represents a handle to a file or directory.
// Each handle is represented by an instance of the FileSystemHandle interface.
// https://developer.mozilla.org/en-US/docs/Web/API/FileSystemHandle
type FileSystemHandle struct {
	js.Value
}

// Properties

// https://developer.mozilla.org/en-US/docs/Web/API/FileSystemHandle/kind
func (h *FileSystemHandle) Kind() string {
	return h.Get("kind").String()
}

// https://developer.mozilla.org/en-US/docs/Web/API/FileSystemHandle/name
func (h *FileSystemHandle) Name() string {
	return h.Get("name").String()
}

// Methods

// https://developer.mozilla.org/en-US/docs/Web/API/FileSystemHandle/isSameEntry
func (h *FileSystemHandle) IsSameEntry(other *FileSystemHandle) *Promise {
	return &Promise{Value: h.Call("isSameEntry", other.Value)}
}

// https://developer.mozilla.org/en-US/docs/Web/API/FileSystemHandle/queryPermission
// options can be 'read', 'readwrite', or 'readwrite-restore'
func (h *FileSystemHandle) QueryPermission(mode string) *Promise {
	opts := map[string]interface{}{"mode": mode}
	return &Promise{Value: h.Call("queryPermission", ToJSValue(opts))}
}

// https://developer.mozilla.org/en-US/docs/Web/API/FileSystemHandle/requestPermission
// options can be 'read', 'readwrite', or 'readwrite-restore'
func (h *FileSystemHandle) RequestPermission(mode string) *Promise {
	opts := map[string]interface{}{"mode": mode}
	return &Promise{Value: h.Call("requestPermission", ToJSValue(opts))}
}

// FileSystemFileHandle provides a handle to a file entry.
// https://developer.mozilla.org/en-US/docs/Web/API/FileSystemFileHandle
type FileSystemFileHandle struct {
	FileSystemHandle
}

// Methods

// https://developer.mozilla.org/en-US/docs/Web/API/FileSystemFileHandle/getFile
func (h *FileSystemFileHandle) GetFile() *Promise {
	return &Promise{Value: h.Call("getFile")}
}

// https://developer.mozilla.org/en-US/docs/Web/API/FileSystemFileHandle/createWritable
// options can include keepExistingData (bool)
func (h *FileSystemFileHandle) CreateWritable(options ...map[string]interface{}) *Promise {
	if len(options) > 0 {
		return &Promise{Value: h.Call("createWritable", ToJSValue(options[0]))}
	}
	return &Promise{Value: h.Call("createWritable")}
}

// https://developer.mozilla.org/en-US/docs/Web/API/FileSystemFileHandle/createSyncAccessHandle
func (h *FileSystemFileHandle) CreateSyncAccessHandle() *Promise {
	return &Promise{Value: h.Call("createSyncAccessHandle")}
}

// FileSystemDirectoryHandle provides a handle to a file system directory.
// https://developer.mozilla.org/en-US/docs/Web/API/FileSystemDirectoryHandle
type FileSystemDirectoryHandle struct {
	FileSystemHandle
}

// Properties

// https://developer.mozilla.org/en-US/docs/Web/API/FileSystemDirectoryHandle/entries
func (h *FileSystemDirectoryHandle) Entries() *AsyncIterable {
	return &AsyncIterable{h.Call("entries")}
}

// Methods

// https://developer.mozilla.org/en-US/docs/Web/API/FileSystemDirectoryHandle/getDirectoryHandle
// options can include create (bool)
func (h *FileSystemDirectoryHandle) GetDirectoryHandle(name string, options ...map[string]interface{}) *Promise {
	if len(options) > 0 {
		return &Promise{Value: h.Call("getDirectoryHandle", name, ToJSValue(options[0]))}
	}
	return &Promise{Value: h.Call("getDirectoryHandle", name)}
}

// https://developer.mozilla.org/en-US/docs/Web/API/FileSystemDirectoryHandle/getFileHandle
// options can include create (bool)
func (h *FileSystemDirectoryHandle) GetFileHandle(name string, options ...map[string]interface{}) *Promise {
	if len(options) > 0 {
		return &Promise{Value: h.Call("getFileHandle", name, ToJSValue(options[0]))}
	}
	return &Promise{Value: h.Call("getFileHandle", name)}
}

// https://developer.mozilla.org/en-US/docs/Web/API/FileSystemDirectoryHandle/removeEntry
// options can include recursive (bool)
func (h *FileSystemDirectoryHandle) RemoveEntry(name string, options ...map[string]interface{}) *Promise {
	if len(options) > 0 {
		return &Promise{Value: h.Call("removeEntry", name, ToJSValue(options[0]))}
	}
	return &Promise{Value: h.Call("removeEntry", name)}
}

// https://developer.mozilla.org/en-US/docs/Web/API/FileSystemDirectoryHandle/resolve
func (h *FileSystemDirectoryHandle) Resolve(possibleDescendant *FileSystemHandle) *Promise {
	return &Promise{Value: h.Call("resolve", possibleDescendant.Value)}
}

// AsyncIterable represents an async iterable, typically returned by entries()
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Iteration_protocols
type AsyncIterable struct {
	js.Value
}

// ToSlice converts an async iterable to a slice of key-value pairs
// This is a convenience method that awaits all entries
func (ai *AsyncIterable) ToSlice() *Promise {
	// Create a JS function to convert async iterable to array
	jsCode := `
		(async (iter) => {
			const results = [];
			for await (const [key, value] of iter) {
				results.push({key: key, value: value});
			}
			return results;
		})
	`
	fn := js.Global().Call("eval", jsCode)
	result := fn.Invoke(ai.Value)
	return &Promise{Value: result}
}

// FileSystemWritableFileStream represents a writable stream to a file.
// https://developer.mozilla.org/en-US/docs/Web/API/FileSystemWritableFileStream
type FileSystemWritableFileStream struct {
	js.Value
}

// Methods

// https://developer.mozilla.org/en-US/docs/Web/API/FileSystemWritableFileStream/write
// data can be a string, Blob, ArrayBuffer, or DataView
func (w *FileSystemWritableFileStream) Write(data interface{}) *Promise {
	return &Promise{Value: w.Call("write", ToJSValue(data))}
}

// https://developer.mozilla.org/en-US/docs/Web/API/FileSystemWritableFileStream/seek
func (w *FileSystemWritableFileStream) Seek(position int) *Promise {
	return &Promise{Value: w.Call("seek", position)}
}

// https://developer.mozilla.org/en-US/docs/Web/API/FileSystemWritableFileStream/truncate
func (w *FileSystemWritableFileStream) Truncate(size int) *Promise {
	return &Promise{Value: w.Call("truncate", size)}
}

// https://developer.mozilla.org/en-US/docs/Web/API/FileSystemWritableFileStream/close
func (w *FileSystemWritableFileStream) Close() *Promise {
	return &Promise{Value: w.Call("close")}
}

// FileSystemDirectoryPickerOptions represents options for directory picker.
// https://developer.mozilla.org/en-US/docs/Web/API/Window/showDirectoryPicker
type FileSystemDirectoryPickerOptions struct {
	// ID is an identifier for the previous directory used
	ID string `json:"id,omitempty"`
	// Mode can be 'read' or 'readwrite'
	Mode string `json:"mode,omitempty"`
	// StartIn specifies the starting directory
	StartIn string `json:"startIn,omitempty"`
}

// FileSystemFilePickerOptions represents options for file picker.
// https://developer.mozilla.org/en-US/docs/Web/API/Window/showOpenFilePicker
type FileSystemFilePickerOptions struct {
	// ExcludeAcceptAllOption
	ExcludeAcceptAllOption bool `json:"excludeAcceptAllOption,omitempty"`
	// ID is an identifier for the previous directory used
	ID string `json:"id,omitempty"`
	// Multiple allows selecting multiple files
	Multiple bool `json:"multiple,omitempty"`
	// Types is an array of accepted file types
	Types []FileType `json:"types,omitempty"`
	// StartIn specifies the starting directory
	StartIn string `json:"startIn,omitempty"`
}

// SaveFilePickerOptions represents options for save file picker.
// https://developer.mozilla.org/en-US/docs/Web/API/Window/showSaveFilePicker
type SaveFilePickerOptions struct {
	// ExcludeAcceptAllOption
	ExcludeAcceptAllOption bool `json:"excludeAcceptAllOption,omitempty"`
	// ID is an identifier for the previous directory used
	ID string `json:"id,omitempty"`
	// SuggestedName provides a suggested file name
	SuggestedName string `json:"suggestedName,omitempty"`
	// Types is an array of accepted file types
	Types []FileType `json:"types,omitempty"`
	// StartIn specifies the starting directory
	StartIn string `json:"startIn,omitempty"`
}

// FileType represents a file type description.
// https://developer.mozilla.org/en-US/docs/Web/API/Window/showOpenFilePicker#types
type FileType struct {
	// Description of the file type
	Description string `json:"description,omitempty"`
	// Accept maps MIME types to file extensions
	Accept map[string][]string `json:"accept"`
}

// Window methods for File System Access API

// ShowOpenFilePicker shows a file picker that allows a user to select a file or files.
// https://developer.mozilla.org/en-US/docs/Web/API/Window/showOpenFilePicker
func (w *window) ShowOpenFilePicker(options ...FileSystemFilePickerOptions) *Promise {
	if len(options) > 0 {
		return &Promise{Value: w.Call("showOpenFilePicker", ToJSValue(options[0]))}
	}
	return &Promise{Value: w.Call("showOpenFilePicker")}
}

// ShowSaveFilePicker shows a file picker that allows a user to save a file.
// https://developer.mozilla.org/en-US/docs/Web/API/Window/showSaveFilePicker
func (w *window) ShowSaveFilePicker(options ...SaveFilePickerOptions) *Promise {
	if len(options) > 0 {
		return &Promise{Value: w.Call("showSaveFilePicker", ToJSValue(options[0]))}
	}
	return &Promise{Value: w.Call("showSaveFilePicker")}
}

// ShowDirectoryPicker shows a directory picker that allows a user to select a directory.
// https://developer.mozilla.org/en-US/docs/Web/API/Window/showDirectoryPicker
func (w *window) ShowDirectoryPicker(options ...FileSystemDirectoryPickerOptions) *Promise {
	if len(options) > 0 {
		return &Promise{Value: w.Call("showDirectoryPicker", ToJSValue(options[0]))}
	}
	return &Promise{Value: w.Call("showDirectoryPicker")}
}

// GetOriginPrivateFileSystem returns a handle for a storage endpoint specific to the origin of the calling code.
// https://developer.mozilla.org/en-US/docs/Web/API/Window/getOriginPrivateFileSystem
func (w *window) GetOriginPrivateFileSystem() *FileSystemDirectoryHandle {
	return &FileSystemDirectoryHandle{FileSystemHandle{w.Call("getOriginPrivateFileSystem")}}
}
