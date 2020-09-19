// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

import fmt "fmt"

func (dst *Picture) SetFields(src *Picture, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "embedded":
			if len(subs) > 0 {
				var newDst, newSrc *Picture_Embedded
				if (src == nil || src.Embedded == nil) && dst.Embedded == nil {
					continue
				}
				if src != nil {
					newSrc = src.Embedded
				}
				if dst.Embedded != nil {
					newDst = dst.Embedded
				} else {
					newDst = &Picture_Embedded{}
					dst.Embedded = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Embedded = src.Embedded
				} else {
					dst.Embedded = nil
				}
			}
		case "sizes":
			if len(subs) > 0 {
				return fmt.Errorf("'sizes' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Sizes = src.Sizes
			} else {
				dst.Sizes = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *Picture_Embedded) SetFields(src *Picture_Embedded, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "mime_type":
			if len(subs) > 0 {
				return fmt.Errorf("'mime_type' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.MimeType = src.MimeType
			} else {
				var zero string
				dst.MimeType = zero
			}
		case "data":
			if len(subs) > 0 {
				return fmt.Errorf("'data' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Data = src.Data
			} else {
				dst.Data = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}
