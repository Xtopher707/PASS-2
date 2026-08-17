Algorithm Sentence_Statistics



Variables

    ch : Character

    length, words, vowels : Integer
Begin
    length := 0
    words := 1
    vowels := 0

    Read(ch)

    While ch <> '.' Do
        length := length + 1

        If ch = ' ' Then
            words := words + 1
        End_If

        If ch = 'a' Or ch = 'e' Or ch = 'i' Or ch = 'o' Or ch = 'u'
           Or ch = 'A' Or ch = 'E' Or ch = 'I' Or ch = 'O' Or ch = 'U' Then
            vowels := vowels + 1
        End_If

        Read(ch)
    End_While

    length := length + 1    // Count the final point

    Write("Length = ", length)
    Write("Words = ", words)
    Write("Vowels = ", vowels)

End