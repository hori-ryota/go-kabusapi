package main

func (doc KabusAPIDocument) TypedArrays() []DefinitionDef {
	typedArrays := make([]DefinitionDef, 0, len(doc.Definitions))
	for _, d := range doc.Definitions {
		if _, ok := d.Type.(ArrayDef); ok {
			typedArrays = append(typedArrays, d)
		}
	}
	return typedArrays
}

func (doc KabusAPIDocument) TypedStructs() []StructDef {
	typedStructs := make([]StructDef, 0, len(doc.Definitions)*2)
	for _, d := range doc.Definitions {
		if a, ok := d.Type.(ArrayDef); ok {
			if ss, ok := a.Elem.(StructDef); ok {
				typedStructs = append(typedStructs, ss)
				for _, p := range ss.Properties {
					if ss, ok := p.Type.(StructDef); ok {
						typedStructs = append(typedStructs, ss)
					}
					if a, ok := p.Type.(ArrayDef); ok {
						if ss, ok := a.Elem.(StructDef); ok {
							typedStructs = append(typedStructs, ss)
						}
					}
				}
				continue
			}
			continue
		}
		if s, ok := d.Type.(StructDef); ok {
			typedStructs = append(typedStructs, s)
			for _, p := range s.Properties {
				if ss, ok := p.Type.(StructDef); ok {
					typedStructs = append(typedStructs, ss)
				}
				if a, ok := p.Type.(ArrayDef); ok {
					if ss, ok := a.Elem.(StructDef); ok {
						typedStructs = append(typedStructs, ss)
					}
				}
			}
		}
	}
	return typedStructs
}
