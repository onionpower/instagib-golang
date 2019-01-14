package main

import "fmt"

type stack struct {
	s  []uint8
	it int
}

func (s *stack) push(val uint8) {
	s.it++
	s.s[s.it] = val
}

func (s *stack) peek() (bool, uint8) {
	if s.it >= len(s.s) || s.it == -1 || s.s[s.it] == 0 {
		return false, 0
	}

	return true, s.s[s.it]
}

func (s *stack) pop() uint8 {
	poped_ix := s.it
	s.it--
	return s.s[poped_ix]
}

const opening_brace = ' '

func isValid(s string) bool {
	s_len := len(s)
	if s_len == 0 {
		return true
	}
	if s_len == 1 {
		return false
	}

	m := makeMap()
	st := &stack{make([]uint8, s_len), -1}
	for i := 0; i < s_len; i++ {
		cur_brace := s[i]
		open_pair, is_brace := m[cur_brace]
		// not a brace -> exit
		if !is_brace {
			return false
		}
		// discover closing brace for opening later
		if open_pair == opening_brace {
			st.push(cur_brace)
			continue
		}
		// match last opened brace with current's opening pair
		ok, last_opened := st.peek()
		if !ok || last_opened != open_pair {
			return false
		}
		st.pop()
	}

	has_content, _ := st.peek()
	return !has_content
}

func makeMap() map[uint8]uint8 {
	m := make(map[uint8]uint8, 6)
	m['('] = opening_brace
	m['['] = opening_brace
	m['{'] = opening_brace
	m[')'] = '('
	m[']'] = '['
	m['}'] = '{'
	return m
}

func main() {
	fmt.Println(isValid("{(}[()]"))
	fmt.Println(isValid("{}[()]"))
	fmt.Println(isValid("}{"))
	fmt.Println(isValid("(("))
}
