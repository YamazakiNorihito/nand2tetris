package compilation_engine

import (
	"bytes"
	"encoding/xml"
	"ny/nand2tetris/compiler1/internal/tokens"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompileClass(t *testing.T) {
	t.Run("should return an error when reader.Read returns an error", func(t *testing.T) {
		/*
			class MainTest {
			}
		*/
		// Arrange
		inputTokensXml := `<tokens>
  <keyword> class</keyword>
  <identifier> MainTest</identifier>
  <symbol> {</symbol>
  <symbol> }</symbol>
</tokens>
`

		expectedXml := `<class>
  <keyword>class</keyword>
  <identifier>MainTest</identifier>
  <symbol>{</symbol>
  <symbol>}</symbol>
</class>`

		reader := strings.NewReader(inputTokensXml)
		testTokens := tokens.LoadTokensFromXML(reader)

		var buf bytes.Buffer
		xmlEnc := xml.NewEncoder(&buf)

		// Act
		ce, err := New(testTokens, xmlEnc)
		assert.NoError(t, err)
		err = ce.CompileClass()

		// Assert
		assert.NoError(t, err)

		assert.Equal(t, expectedXml, buf.String())
	})

	t.Run("should return correct XML for ArrayTest class", func(t *testing.T) {
		/*
			class MainTest {
			}
		*/
		// Arrange
		inputTokensXml := `<tokens>
  <keyword> class </keyword>
  <identifier> Main </identifier>
  <symbol> { </symbol>
  <keyword> function </keyword>
  <keyword> void </keyword>
  <identifier> main </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> var </keyword>
  <identifier> Array </identifier>
  <identifier> a </identifier>
  <symbol> ; </symbol>
  <keyword> var </keyword>
  <keyword> int </keyword>
  <identifier> length </identifier>
  <symbol> ; </symbol>
  <keyword> var </keyword>
  <keyword> int </keyword>
  <identifier> i </identifier>
  <symbol> , </symbol>
  <identifier> sum </identifier>
  <symbol> ; </symbol>
  <keyword> let </keyword>
  <identifier> length </identifier>
  <symbol> = </symbol>
  <identifier> Keyboard </identifier>
  <symbol> . </symbol>
  <identifier> readInt </identifier>
  <symbol> ( </symbol>
  <stringConstant> HOW MANY NUMBERS?  </stringConstant>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> let </keyword>
  <identifier> a </identifier>
  <symbol> = </symbol>
  <identifier> Array </identifier>
  <symbol> . </symbol>
  <identifier> new </identifier>
  <symbol> ( </symbol>
  <identifier> length </identifier>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> let </keyword>
  <identifier> i </identifier>
  <symbol> = </symbol>
  <integerConstant> 0 </integerConstant>
  <symbol> ; </symbol>
  <keyword> while </keyword>
  <symbol> ( </symbol>
  <identifier> i </identifier>
  <symbol> &lt; </symbol>
  <identifier> length </identifier>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> let </keyword>
  <identifier> a </identifier>
  <symbol> [ </symbol>
  <identifier> i </identifier>
  <symbol> ] </symbol>
  <symbol> = </symbol>
  <identifier> Keyboard </identifier>
  <symbol> . </symbol>
  <identifier> readInt </identifier>
  <symbol> ( </symbol>
  <stringConstant> ENTER THE NEXT NUMBER:  </stringConstant>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> let </keyword>
  <identifier> i </identifier>
  <symbol> = </symbol>
  <identifier> i </identifier>
  <symbol> + </symbol>
  <integerConstant> 1 </integerConstant>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> let </keyword>
  <identifier> i </identifier>
  <symbol> = </symbol>
  <integerConstant> 0 </integerConstant>
  <symbol> ; </symbol>
  <keyword> let </keyword>
  <identifier> sum </identifier>
  <symbol> = </symbol>
  <integerConstant> 0 </integerConstant>
  <symbol> ; </symbol>
  <keyword> while </keyword>
  <symbol> ( </symbol>
  <identifier> i </identifier>
  <symbol> &lt; </symbol>
  <identifier> length </identifier>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> let </keyword>
  <identifier> sum </identifier>
  <symbol> = </symbol>
  <identifier> sum </identifier>
  <symbol> + </symbol>
  <identifier> a </identifier>
  <symbol> [ </symbol>
  <identifier> i </identifier>
  <symbol> ] </symbol>
  <symbol> ; </symbol>
  <keyword> let </keyword>
  <identifier> i </identifier>
  <symbol> = </symbol>
  <identifier> i </identifier>
  <symbol> + </symbol>
  <integerConstant> 1 </integerConstant>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> do </keyword>
  <identifier> Output </identifier>
  <symbol> . </symbol>
  <identifier> printString </identifier>
  <symbol> ( </symbol>
  <stringConstant> THE AVERAGE IS:  </stringConstant>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> do </keyword>
  <identifier> Output </identifier>
  <symbol> . </symbol>
  <identifier> printInt </identifier>
  <symbol> ( </symbol>
  <identifier> sum </identifier>
  <symbol> / </symbol>
  <identifier> length </identifier>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> do </keyword>
  <identifier> Output </identifier>
  <symbol> . </symbol>
  <identifier> println </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> return </keyword>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <symbol> } </symbol>
</tokens>
`

		expectedXml := `<class>
  <keyword>class</keyword>
  <identifier>Main</identifier>
  <symbol>{</symbol>
  <subroutineDec>
    <keyword>function</keyword>
    <keyword>void</keyword>
    <identifier>main</identifier>
    <symbol>(</symbol>
    <parameterList></parameterList>
    <symbol>)</symbol>
    <subroutineBody>
      <symbol>{</symbol>
      <varDec>
        <keyword>var</keyword>
        <identifier>Array</identifier>
        <identifier>a</identifier>
        <symbol>;</symbol>
      </varDec>
      <varDec>
        <keyword>var</keyword>
        <keyword>int</keyword>
        <identifier>length</identifier>
        <symbol>;</symbol>
      </varDec>
      <varDec>
        <keyword>var</keyword>
        <keyword>int</keyword>
        <identifier>i</identifier>
        <symbol>,</symbol>
        <identifier>sum</identifier>
        <symbol>;</symbol>
      </varDec>
      <statements>
        <letStatement>
          <keyword>let</keyword>
          <identifier>length</identifier>
          <symbol>=</symbol>
          <expression>
            <term>
              <identifier>Keyboard</identifier>
              <symbol>.</symbol>
              <identifier>readInt</identifier>
              <symbol>(</symbol>
              <expressionList>
                <expression>
                  <term>
                    <stringConstant>HOW MANY NUMBERS? </stringConstant>
                  </term>
                </expression>
              </expressionList>
              <symbol>)</symbol>
            </term>
          </expression>
          <symbol>;</symbol>
        </letStatement>
        <letStatement>
          <keyword>let</keyword>
          <identifier>a</identifier>
          <symbol>=</symbol>
          <expression>
            <term>
              <identifier>Array</identifier>
              <symbol>.</symbol>
              <identifier>new</identifier>
              <symbol>(</symbol>
              <expressionList>
                <expression>
                  <term>
                    <identifier>length</identifier>
                  </term>
                </expression>
              </expressionList>
              <symbol>)</symbol>
            </term>
          </expression>
          <symbol>;</symbol>
        </letStatement>
        <letStatement>
          <keyword>let</keyword>
          <identifier>i</identifier>
          <symbol>=</symbol>
          <expression>
            <term>
              <integerConstant>0</integerConstant>
            </term>
          </expression>
          <symbol>;</symbol>
        </letStatement>
        <whileStatement>
          <keyword>while</keyword>
          <symbol>(</symbol>
          <expression>
            <term>
              <identifier>i</identifier>
            </term>
            <symbol>&lt;</symbol>
            <term>
              <identifier>length</identifier>
            </term>
          </expression>
          <symbol>)</symbol>
          <symbol>{</symbol>
          <statements>
            <letStatement>
              <keyword>let</keyword>
              <identifier>a</identifier>
              <symbol>[</symbol>
              <expression>
                <term>
                  <identifier>i</identifier>
                </term>
              </expression>
              <symbol>]</symbol>
              <symbol>=</symbol>
              <expression>
                <term>
                  <identifier>Keyboard</identifier>
                  <symbol>.</symbol>
                  <identifier>readInt</identifier>
                  <symbol>(</symbol>
                  <expressionList>
                    <expression>
                      <term>
                        <stringConstant>ENTER THE NEXT NUMBER: </stringConstant>
                      </term>
                    </expression>
                  </expressionList>
                  <symbol>)</symbol>
                </term>
              </expression>
              <symbol>;</symbol>
            </letStatement>
            <letStatement>
              <keyword>let</keyword>
              <identifier>i</identifier>
              <symbol>=</symbol>
              <expression>
                <term>
                  <identifier>i</identifier>
                </term>
                <symbol>+</symbol>
                <term>
                  <integerConstant>1</integerConstant>
                </term>
              </expression>
              <symbol>;</symbol>
            </letStatement>
          </statements>
          <symbol>}</symbol>
        </whileStatement>
        <letStatement>
          <keyword>let</keyword>
          <identifier>i</identifier>
          <symbol>=</symbol>
          <expression>
            <term>
              <integerConstant>0</integerConstant>
            </term>
          </expression>
          <symbol>;</symbol>
        </letStatement>
        <letStatement>
          <keyword>let</keyword>
          <identifier>sum</identifier>
          <symbol>=</symbol>
          <expression>
            <term>
              <integerConstant>0</integerConstant>
            </term>
          </expression>
          <symbol>;</symbol>
        </letStatement>
        <whileStatement>
          <keyword>while</keyword>
          <symbol>(</symbol>
          <expression>
            <term>
              <identifier>i</identifier>
            </term>
            <symbol>&lt;</symbol>
            <term>
              <identifier>length</identifier>
            </term>
          </expression>
          <symbol>)</symbol>
          <symbol>{</symbol>
          <statements>
            <letStatement>
              <keyword>let</keyword>
              <identifier>sum</identifier>
              <symbol>=</symbol>
              <expression>
                <term>
                  <identifier>sum</identifier>
                </term>
                <symbol>+</symbol>
                <term>
                  <identifier>a</identifier>
                  <symbol>[</symbol>
                  <expression>
                    <term>
                      <identifier>i</identifier>
                    </term>
                  </expression>
                  <symbol>]</symbol>
                </term>
              </expression>
              <symbol>;</symbol>
            </letStatement>
            <letStatement>
              <keyword>let</keyword>
              <identifier>i</identifier>
              <symbol>=</symbol>
              <expression>
                <term>
                  <identifier>i</identifier>
                </term>
                <symbol>+</symbol>
                <term>
                  <integerConstant>1</integerConstant>
                </term>
              </expression>
              <symbol>;</symbol>
            </letStatement>
          </statements>
          <symbol>}</symbol>
        </whileStatement>
        <doStatement>
          <keyword>do</keyword>
          <identifier>Output</identifier>
          <symbol>.</symbol>
          <identifier>printString</identifier>
          <symbol>(</symbol>
          <expressionList>
            <expression>
              <term>
                <stringConstant>THE AVERAGE IS: </stringConstant>
              </term>
            </expression>
          </expressionList>
          <symbol>)</symbol>
          <symbol>;</symbol>
        </doStatement>
        <doStatement>
          <keyword>do</keyword>
          <identifier>Output</identifier>
          <symbol>.</symbol>
          <identifier>printInt</identifier>
          <symbol>(</symbol>
          <expressionList>
            <expression>
              <term>
                <identifier>sum</identifier>
              </term>
              <symbol>/</symbol>
              <term>
                <identifier>length</identifier>
              </term>
            </expression>
          </expressionList>
          <symbol>)</symbol>
          <symbol>;</symbol>
        </doStatement>
        <doStatement>
          <keyword>do</keyword>
          <identifier>Output</identifier>
          <symbol>.</symbol>
          <identifier>println</identifier>
          <symbol>(</symbol>
          <expressionList></expressionList>
          <symbol>)</symbol>
          <symbol>;</symbol>
        </doStatement>
        <returnStatement>
          <keyword>return</keyword>
          <symbol>;</symbol>
        </returnStatement>
      </statements>
      <symbol>}</symbol>
    </subroutineBody>
  </subroutineDec>
  <symbol>}</symbol>
</class>`

		reader := strings.NewReader(inputTokensXml)
		testTokens := tokens.LoadTokensFromXML(reader)

		var buf bytes.Buffer
		xmlEnc := xml.NewEncoder(&buf)

		// Act
		ce, err := New(testTokens, xmlEnc)
		assert.NoError(t, err)
		err = ce.CompileClass()

		// Assert
		assert.NoError(t, err)

		assert.Equal(t, expectedXml, buf.String())
	})

	t.Run("should return correct XML for Square class", func(t *testing.T) {
		t.Run("should return XML when compiling Main.jack", func(t *testing.T) {

			// Arrange
			inputTokensXml := `<tokens>
  <keyword> class </keyword>
  <identifier> Main </identifier>
  <symbol> { </symbol>
  <keyword> static </keyword>
  <keyword> boolean </keyword>
  <identifier> test </identifier>
  <symbol> ; </symbol>
  <keyword> function </keyword>
  <keyword> void </keyword>
  <identifier> main </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> var </keyword>
  <identifier> SquareGame </identifier>
  <identifier> game </identifier>
  <symbol> ; </symbol>
  <keyword> let </keyword>
  <identifier> game </identifier>
  <symbol> = </symbol>
  <identifier> SquareGame </identifier>
  <symbol> . </symbol>
  <identifier> new </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> do </keyword>
  <identifier> game </identifier>
  <symbol> . </symbol>
  <identifier> run </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> do </keyword>
  <identifier> game </identifier>
  <symbol> . </symbol>
  <identifier> dispose </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> return </keyword>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> function </keyword>
  <keyword> void </keyword>
  <identifier> more </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> var </keyword>
  <keyword> int </keyword>
  <identifier> i </identifier>
  <symbol> , </symbol>
  <identifier> j </identifier>
  <symbol> ; </symbol>
  <keyword> var </keyword>
  <identifier> String </identifier>
  <identifier> s </identifier>
  <symbol> ; </symbol>
  <keyword> var </keyword>
  <identifier> Array </identifier>
  <identifier> a </identifier>
  <symbol> ; </symbol>
  <keyword> if </keyword>
  <symbol> ( </symbol>
  <keyword> false </keyword>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> let </keyword>
  <identifier> s </identifier>
  <symbol> = </symbol>
  <stringConstant> string constant </stringConstant>
  <symbol> ; </symbol>
  <keyword> let </keyword>
  <identifier> s </identifier>
  <symbol> = </symbol>
  <keyword> null </keyword>
  <symbol> ; </symbol>
  <keyword> let </keyword>
  <identifier> a </identifier>
  <symbol> [ </symbol>
  <integerConstant> 1 </integerConstant>
  <symbol> ] </symbol>
  <symbol> = </symbol>
  <identifier> a </identifier>
  <symbol> [ </symbol>
  <integerConstant> 2 </integerConstant>
  <symbol> ] </symbol>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> else </keyword>
  <symbol> { </symbol>
  <keyword> let </keyword>
  <identifier> i </identifier>
  <symbol> = </symbol>
  <identifier> i </identifier>
  <symbol> * </symbol>
  <symbol> ( </symbol>
  <symbol> - </symbol>
  <identifier> j </identifier>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> let </keyword>
  <identifier> j </identifier>
  <symbol> = </symbol>
  <identifier> j </identifier>
  <symbol> / </symbol>
  <symbol> ( </symbol>
  <symbol> - </symbol>
  <integerConstant> 2 </integerConstant>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> let </keyword>
  <identifier> i </identifier>
  <symbol> = </symbol>
  <identifier> i </identifier>
  <symbol> | </symbol>
  <identifier> j </identifier>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> return </keyword>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <symbol> } </symbol>
</tokens>
`

			expectedXml := `<class>
  <keyword>class</keyword>
  <identifier>Main</identifier>
  <symbol>{</symbol>
  <classVarDec>
    <keyword>static</keyword>
    <keyword>boolean</keyword>
    <identifier>test</identifier>
    <symbol>;</symbol>
  </classVarDec>
  <subroutineDec>
    <keyword>function</keyword>
    <keyword>void</keyword>
    <identifier>main</identifier>
    <symbol>(</symbol>
    <parameterList></parameterList>
    <symbol>)</symbol>
    <subroutineBody>
      <symbol>{</symbol>
      <varDec>
        <keyword>var</keyword>
        <identifier>SquareGame</identifier>
        <identifier>game</identifier>
        <symbol>;</symbol>
      </varDec>
      <statements>
        <letStatement>
          <keyword>let</keyword>
          <identifier>game</identifier>
          <symbol>=</symbol>
          <expression>
            <term>
              <identifier>SquareGame</identifier>
              <symbol>.</symbol>
              <identifier>new</identifier>
              <symbol>(</symbol>
              <expressionList></expressionList>
              <symbol>)</symbol>
            </term>
          </expression>
          <symbol>;</symbol>
        </letStatement>
        <doStatement>
          <keyword>do</keyword>
          <identifier>game</identifier>
          <symbol>.</symbol>
          <identifier>run</identifier>
          <symbol>(</symbol>
          <expressionList></expressionList>
          <symbol>)</symbol>
          <symbol>;</symbol>
        </doStatement>
        <doStatement>
          <keyword>do</keyword>
          <identifier>game</identifier>
          <symbol>.</symbol>
          <identifier>dispose</identifier>
          <symbol>(</symbol>
          <expressionList></expressionList>
          <symbol>)</symbol>
          <symbol>;</symbol>
        </doStatement>
        <returnStatement>
          <keyword>return</keyword>
          <symbol>;</symbol>
        </returnStatement>
      </statements>
      <symbol>}</symbol>
    </subroutineBody>
  </subroutineDec>
  <subroutineDec>
    <keyword>function</keyword>
    <keyword>void</keyword>
    <identifier>more</identifier>
    <symbol>(</symbol>
    <parameterList></parameterList>
    <symbol>)</symbol>
    <subroutineBody>
      <symbol>{</symbol>
      <varDec>
        <keyword>var</keyword>
        <keyword>int</keyword>
        <identifier>i</identifier>
        <symbol>,</symbol>
        <identifier>j</identifier>
        <symbol>;</symbol>
      </varDec>
      <varDec>
        <keyword>var</keyword>
        <identifier>String</identifier>
        <identifier>s</identifier>
        <symbol>;</symbol>
      </varDec>
      <varDec>
        <keyword>var</keyword>
        <identifier>Array</identifier>
        <identifier>a</identifier>
        <symbol>;</symbol>
      </varDec>
      <statements>
        <ifStatement>
          <keyword>if</keyword>
          <symbol>(</symbol>
          <expression>
            <term>
              <keyword>false</keyword>
            </term>
          </expression>
          <symbol>)</symbol>
          <symbol>{</symbol>
          <statements>
            <letStatement>
              <keyword>let</keyword>
              <identifier>s</identifier>
              <symbol>=</symbol>
              <expression>
                <term>
                  <stringConstant>string constant</stringConstant>
                </term>
              </expression>
              <symbol>;</symbol>
            </letStatement>
            <letStatement>
              <keyword>let</keyword>
              <identifier>s</identifier>
              <symbol>=</symbol>
              <expression>
                <term>
                  <keyword>null</keyword>
                </term>
              </expression>
              <symbol>;</symbol>
            </letStatement>
            <letStatement>
              <keyword>let</keyword>
              <identifier>a</identifier>
              <symbol>[</symbol>
              <expression>
                <term>
                  <integerConstant>1</integerConstant>
                </term>
              </expression>
              <symbol>]</symbol>
              <symbol>=</symbol>
              <expression>
                <term>
                  <identifier>a</identifier>
                  <symbol>[</symbol>
                  <expression>
                    <term>
                      <integerConstant>2</integerConstant>
                    </term>
                  </expression>
                  <symbol>]</symbol>
                </term>
              </expression>
              <symbol>;</symbol>
            </letStatement>
          </statements>
          <symbol>}</symbol>
          <keyword>else</keyword>
          <symbol>{</symbol>
          <statements>
            <letStatement>
              <keyword>let</keyword>
              <identifier>i</identifier>
              <symbol>=</symbol>
              <expression>
                <term>
                  <identifier>i</identifier>
                </term>
                <symbol>*</symbol>
                <term>
                  <symbol>(</symbol>
                  <expression>
                    <term>
                      <symbol>-</symbol>
                      <term>
                        <identifier>j</identifier>
                      </term>
                    </term>
                  </expression>
                  <symbol>)</symbol>
                </term>
              </expression>
              <symbol>;</symbol>
            </letStatement>
            <letStatement>
              <keyword>let</keyword>
              <identifier>j</identifier>
              <symbol>=</symbol>
              <expression>
                <term>
                  <identifier>j</identifier>
                </term>
                <symbol>/</symbol>
                <term>
                  <symbol>(</symbol>
                  <expression>
                    <term>
                      <symbol>-</symbol>
                      <term>
                        <integerConstant>2</integerConstant>
                      </term>
                    </term>
                  </expression>
                  <symbol>)</symbol>
                </term>
              </expression>
              <symbol>;</symbol>
            </letStatement>
            <letStatement>
              <keyword>let</keyword>
              <identifier>i</identifier>
              <symbol>=</symbol>
              <expression>
                <term>
                  <identifier>i</identifier>
                </term>
                <symbol>|</symbol>
                <term>
                  <identifier>j</identifier>
                </term>
              </expression>
              <symbol>;</symbol>
            </letStatement>
          </statements>
          <symbol>}</symbol>
        </ifStatement>
        <returnStatement>
          <keyword>return</keyword>
          <symbol>;</symbol>
        </returnStatement>
      </statements>
      <symbol>}</symbol>
    </subroutineBody>
  </subroutineDec>
  <symbol>}</symbol>
</class>`

			reader := strings.NewReader(inputTokensXml)
			testTokens := tokens.LoadTokensFromXML(reader)

			var buf bytes.Buffer
			xmlEnc := xml.NewEncoder(&buf)

			// Act
			ce, err := New(testTokens, xmlEnc)
			assert.NoError(t, err)
			err = ce.CompileClass()

			// Assert
			assert.NoError(t, err)

			assert.Equal(t, expectedXml, buf.String())
		})
		t.Run("should return XML when compiling Square.jack", func(t *testing.T) {

			// Arrange
			inputTokensXml := `<tokens>
  <keyword> class </keyword>
  <identifier> Square </identifier>
  <symbol> { </symbol>
  <keyword> field </keyword>
  <keyword> int </keyword>
  <identifier> x </identifier>
  <symbol> , </symbol>
  <identifier> y </identifier>
  <symbol> ; </symbol>
  <keyword> field </keyword>
  <keyword> int </keyword>
  <identifier> size </identifier>
  <symbol> ; </symbol>
  <keyword> constructor </keyword>
  <identifier> Square </identifier>
  <identifier> new </identifier>
  <symbol> ( </symbol>
  <keyword> int </keyword>
  <identifier> Ax </identifier>
  <symbol> , </symbol>
  <keyword> int </keyword>
  <identifier> Ay </identifier>
  <symbol> , </symbol>
  <keyword> int </keyword>
  <identifier> Asize </identifier>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> let </keyword>
  <identifier> x </identifier>
  <symbol> = </symbol>
  <identifier> Ax </identifier>
  <symbol> ; </symbol>
  <keyword> let </keyword>
  <identifier> y </identifier>
  <symbol> = </symbol>
  <identifier> Ay </identifier>
  <symbol> ; </symbol>
  <keyword> let </keyword>
  <identifier> size </identifier>
  <symbol> = </symbol>
  <identifier> Asize </identifier>
  <symbol> ; </symbol>
  <keyword> do </keyword>
  <identifier> draw </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> return </keyword>
  <keyword> this </keyword>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> method </keyword>
  <keyword> void </keyword>
  <identifier> dispose </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> do </keyword>
  <identifier> Memory </identifier>
  <symbol> . </symbol>
  <identifier> deAlloc </identifier>
  <symbol> ( </symbol>
  <keyword> this </keyword>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> return </keyword>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> method </keyword>
  <keyword> void </keyword>
  <identifier> draw </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> do </keyword>
  <identifier> Screen </identifier>
  <symbol> . </symbol>
  <identifier> setColor </identifier>
  <symbol> ( </symbol>
  <keyword> true </keyword>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> do </keyword>
  <identifier> Screen </identifier>
  <symbol> . </symbol>
  <identifier> drawRectangle </identifier>
  <symbol> ( </symbol>
  <identifier> x </identifier>
  <symbol> , </symbol>
  <identifier> y </identifier>
  <symbol> , </symbol>
  <identifier> x </identifier>
  <symbol> + </symbol>
  <identifier> size </identifier>
  <symbol> , </symbol>
  <identifier> y </identifier>
  <symbol> + </symbol>
  <identifier> size </identifier>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> return </keyword>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> method </keyword>
  <keyword> void </keyword>
  <identifier> erase </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> do </keyword>
  <identifier> Screen </identifier>
  <symbol> . </symbol>
  <identifier> setColor </identifier>
  <symbol> ( </symbol>
  <keyword> false </keyword>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> do </keyword>
  <identifier> Screen </identifier>
  <symbol> . </symbol>
  <identifier> drawRectangle </identifier>
  <symbol> ( </symbol>
  <identifier> x </identifier>
  <symbol> , </symbol>
  <identifier> y </identifier>
  <symbol> , </symbol>
  <identifier> x </identifier>
  <symbol> + </symbol>
  <identifier> size </identifier>
  <symbol> , </symbol>
  <identifier> y </identifier>
  <symbol> + </symbol>
  <identifier> size </identifier>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> return </keyword>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> method </keyword>
  <keyword> void </keyword>
  <identifier> incSize </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> if </keyword>
  <symbol> ( </symbol>
  <symbol> ( </symbol>
  <symbol> ( </symbol>
  <identifier> y </identifier>
  <symbol> + </symbol>
  <identifier> size </identifier>
  <symbol> ) </symbol>
  <symbol> &lt; </symbol>
  <integerConstant> 254 </integerConstant>
  <symbol> ) </symbol>
  <symbol> &amp; </symbol>
  <symbol> ( </symbol>
  <symbol> ( </symbol>
  <identifier> x </identifier>
  <symbol> + </symbol>
  <identifier> size </identifier>
  <symbol> ) </symbol>
  <symbol> &lt; </symbol>
  <integerConstant> 510 </integerConstant>
  <symbol> ) </symbol>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> do </keyword>
  <identifier> erase </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> let </keyword>
  <identifier> size </identifier>
  <symbol> = </symbol>
  <identifier> size </identifier>
  <symbol> + </symbol>
  <integerConstant> 2 </integerConstant>
  <symbol> ; </symbol>
  <keyword> do </keyword>
  <identifier> draw </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> return </keyword>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> method </keyword>
  <keyword> void </keyword>
  <identifier> decSize </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> if </keyword>
  <symbol> ( </symbol>
  <identifier> size </identifier>
  <symbol> &gt; </symbol>
  <integerConstant> 2 </integerConstant>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> do </keyword>
  <identifier> erase </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> let </keyword>
  <identifier> size </identifier>
  <symbol> = </symbol>
  <identifier> size </identifier>
  <symbol> - </symbol>
  <integerConstant> 2 </integerConstant>
  <symbol> ; </symbol>
  <keyword> do </keyword>
  <identifier> draw </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> return </keyword>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> method </keyword>
  <keyword> void </keyword>
  <identifier> moveUp </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> if </keyword>
  <symbol> ( </symbol>
  <identifier> y </identifier>
  <symbol> &gt; </symbol>
  <integerConstant> 1 </integerConstant>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> do </keyword>
  <identifier> Screen </identifier>
  <symbol> . </symbol>
  <identifier> setColor </identifier>
  <symbol> ( </symbol>
  <keyword> false </keyword>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> do </keyword>
  <identifier> Screen </identifier>
  <symbol> . </symbol>
  <identifier> drawRectangle </identifier>
  <symbol> ( </symbol>
  <identifier> x </identifier>
  <symbol> , </symbol>
  <symbol> ( </symbol>
  <identifier> y </identifier>
  <symbol> + </symbol>
  <identifier> size </identifier>
  <symbol> ) </symbol>
  <symbol> - </symbol>
  <integerConstant> 1 </integerConstant>
  <symbol> , </symbol>
  <identifier> x </identifier>
  <symbol> + </symbol>
  <identifier> size </identifier>
  <symbol> , </symbol>
  <identifier> y </identifier>
  <symbol> + </symbol>
  <identifier> size </identifier>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> let </keyword>
  <identifier> y </identifier>
  <symbol> = </symbol>
  <identifier> y </identifier>
  <symbol> - </symbol>
  <integerConstant> 2 </integerConstant>
  <symbol> ; </symbol>
  <keyword> do </keyword>
  <identifier> Screen </identifier>
  <symbol> . </symbol>
  <identifier> setColor </identifier>
  <symbol> ( </symbol>
  <keyword> true </keyword>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> do </keyword>
  <identifier> Screen </identifier>
  <symbol> . </symbol>
  <identifier> drawRectangle </identifier>
  <symbol> ( </symbol>
  <identifier> x </identifier>
  <symbol> , </symbol>
  <identifier> y </identifier>
  <symbol> , </symbol>
  <identifier> x </identifier>
  <symbol> + </symbol>
  <identifier> size </identifier>
  <symbol> , </symbol>
  <identifier> y </identifier>
  <symbol> + </symbol>
  <integerConstant> 1 </integerConstant>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> return </keyword>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> method </keyword>
  <keyword> void </keyword>
  <identifier> moveDown </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> if </keyword>
  <symbol> ( </symbol>
  <symbol> ( </symbol>
  <identifier> y </identifier>
  <symbol> + </symbol>
  <identifier> size </identifier>
  <symbol> ) </symbol>
  <symbol> &lt; </symbol>
  <integerConstant> 254 </integerConstant>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> do </keyword>
  <identifier> Screen </identifier>
  <symbol> . </symbol>
  <identifier> setColor </identifier>
  <symbol> ( </symbol>
  <keyword> false </keyword>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> do </keyword>
  <identifier> Screen </identifier>
  <symbol> . </symbol>
  <identifier> drawRectangle </identifier>
  <symbol> ( </symbol>
  <identifier> x </identifier>
  <symbol> , </symbol>
  <identifier> y </identifier>
  <symbol> , </symbol>
  <identifier> x </identifier>
  <symbol> + </symbol>
  <identifier> size </identifier>
  <symbol> , </symbol>
  <identifier> y </identifier>
  <symbol> + </symbol>
  <integerConstant> 1 </integerConstant>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> let </keyword>
  <identifier> y </identifier>
  <symbol> = </symbol>
  <identifier> y </identifier>
  <symbol> + </symbol>
  <integerConstant> 2 </integerConstant>
  <symbol> ; </symbol>
  <keyword> do </keyword>
  <identifier> Screen </identifier>
  <symbol> . </symbol>
  <identifier> setColor </identifier>
  <symbol> ( </symbol>
  <keyword> true </keyword>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> do </keyword>
  <identifier> Screen </identifier>
  <symbol> . </symbol>
  <identifier> drawRectangle </identifier>
  <symbol> ( </symbol>
  <identifier> x </identifier>
  <symbol> , </symbol>
  <symbol> ( </symbol>
  <identifier> y </identifier>
  <symbol> + </symbol>
  <identifier> size </identifier>
  <symbol> ) </symbol>
  <symbol> - </symbol>
  <integerConstant> 1 </integerConstant>
  <symbol> , </symbol>
  <identifier> x </identifier>
  <symbol> + </symbol>
  <identifier> size </identifier>
  <symbol> , </symbol>
  <identifier> y </identifier>
  <symbol> + </symbol>
  <identifier> size </identifier>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> return </keyword>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> method </keyword>
  <keyword> void </keyword>
  <identifier> moveLeft </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> if </keyword>
  <symbol> ( </symbol>
  <identifier> x </identifier>
  <symbol> &gt; </symbol>
  <integerConstant> 1 </integerConstant>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> do </keyword>
  <identifier> Screen </identifier>
  <symbol> . </symbol>
  <identifier> setColor </identifier>
  <symbol> ( </symbol>
  <keyword> false </keyword>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> do </keyword>
  <identifier> Screen </identifier>
  <symbol> . </symbol>
  <identifier> drawRectangle </identifier>
  <symbol> ( </symbol>
  <symbol> ( </symbol>
  <identifier> x </identifier>
  <symbol> + </symbol>
  <identifier> size </identifier>
  <symbol> ) </symbol>
  <symbol> - </symbol>
  <integerConstant> 1 </integerConstant>
  <symbol> , </symbol>
  <identifier> y </identifier>
  <symbol> , </symbol>
  <identifier> x </identifier>
  <symbol> + </symbol>
  <identifier> size </identifier>
  <symbol> , </symbol>
  <identifier> y </identifier>
  <symbol> + </symbol>
  <identifier> size </identifier>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> let </keyword>
  <identifier> x </identifier>
  <symbol> = </symbol>
  <identifier> x </identifier>
  <symbol> - </symbol>
  <integerConstant> 2 </integerConstant>
  <symbol> ; </symbol>
  <keyword> do </keyword>
  <identifier> Screen </identifier>
  <symbol> . </symbol>
  <identifier> setColor </identifier>
  <symbol> ( </symbol>
  <keyword> true </keyword>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> do </keyword>
  <identifier> Screen </identifier>
  <symbol> . </symbol>
  <identifier> drawRectangle </identifier>
  <symbol> ( </symbol>
  <identifier> x </identifier>
  <symbol> , </symbol>
  <identifier> y </identifier>
  <symbol> , </symbol>
  <identifier> x </identifier>
  <symbol> + </symbol>
  <integerConstant> 1 </integerConstant>
  <symbol> , </symbol>
  <identifier> y </identifier>
  <symbol> + </symbol>
  <identifier> size </identifier>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> return </keyword>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> method </keyword>
  <keyword> void </keyword>
  <identifier> moveRight </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> if </keyword>
  <symbol> ( </symbol>
  <symbol> ( </symbol>
  <identifier> x </identifier>
  <symbol> + </symbol>
  <identifier> size </identifier>
  <symbol> ) </symbol>
  <symbol> &lt; </symbol>
  <integerConstant> 510 </integerConstant>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> do </keyword>
  <identifier> Screen </identifier>
  <symbol> . </symbol>
  <identifier> setColor </identifier>
  <symbol> ( </symbol>
  <keyword> false </keyword>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> do </keyword>
  <identifier> Screen </identifier>
  <symbol> . </symbol>
  <identifier> drawRectangle </identifier>
  <symbol> ( </symbol>
  <identifier> x </identifier>
  <symbol> , </symbol>
  <identifier> y </identifier>
  <symbol> , </symbol>
  <identifier> x </identifier>
  <symbol> + </symbol>
  <integerConstant> 1 </integerConstant>
  <symbol> , </symbol>
  <identifier> y </identifier>
  <symbol> + </symbol>
  <identifier> size </identifier>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> let </keyword>
  <identifier> x </identifier>
  <symbol> = </symbol>
  <identifier> x </identifier>
  <symbol> + </symbol>
  <integerConstant> 2 </integerConstant>
  <symbol> ; </symbol>
  <keyword> do </keyword>
  <identifier> Screen </identifier>
  <symbol> . </symbol>
  <identifier> setColor </identifier>
  <symbol> ( </symbol>
  <keyword> true </keyword>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> do </keyword>
  <identifier> Screen </identifier>
  <symbol> . </symbol>
  <identifier> drawRectangle </identifier>
  <symbol> ( </symbol>
  <symbol> ( </symbol>
  <identifier> x </identifier>
  <symbol> + </symbol>
  <identifier> size </identifier>
  <symbol> ) </symbol>
  <symbol> - </symbol>
  <integerConstant> 1 </integerConstant>
  <symbol> , </symbol>
  <identifier> y </identifier>
  <symbol> , </symbol>
  <identifier> x </identifier>
  <symbol> + </symbol>
  <identifier> size </identifier>
  <symbol> , </symbol>
  <identifier> y </identifier>
  <symbol> + </symbol>
  <identifier> size </identifier>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> return </keyword>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <symbol> } </symbol>
</tokens>
`

			expectedXml := `<class>
  <keyword>class</keyword>
  <identifier>Square</identifier>
  <symbol>{</symbol>
  <classVarDec>
    <keyword>field</keyword>
    <keyword>int</keyword>
    <identifier>x</identifier>
    <symbol>,</symbol>
    <identifier>y</identifier>
    <symbol>;</symbol>
  </classVarDec>
  <classVarDec>
    <keyword>field</keyword>
    <keyword>int</keyword>
    <identifier>size</identifier>
    <symbol>;</symbol>
  </classVarDec>
  <subroutineDec>
    <keyword>constructor</keyword>
    <identifier>Square</identifier>
    <identifier>new</identifier>
    <symbol>(</symbol>
    <parameterList>
      <keyword>int</keyword>
      <identifier>Ax</identifier>
      <symbol>,</symbol>
      <keyword>int</keyword>
      <identifier>Ay</identifier>
      <symbol>,</symbol>
      <keyword>int</keyword>
      <identifier>Asize</identifier>
    </parameterList>
    <symbol>)</symbol>
    <subroutineBody>
      <symbol>{</symbol>
      <statements>
        <letStatement>
          <keyword>let</keyword>
          <identifier>x</identifier>
          <symbol>=</symbol>
          <expression>
            <term>
              <identifier>Ax</identifier>
            </term>
          </expression>
          <symbol>;</symbol>
        </letStatement>
        <letStatement>
          <keyword>let</keyword>
          <identifier>y</identifier>
          <symbol>=</symbol>
          <expression>
            <term>
              <identifier>Ay</identifier>
            </term>
          </expression>
          <symbol>;</symbol>
        </letStatement>
        <letStatement>
          <keyword>let</keyword>
          <identifier>size</identifier>
          <symbol>=</symbol>
          <expression>
            <term>
              <identifier>Asize</identifier>
            </term>
          </expression>
          <symbol>;</symbol>
        </letStatement>
        <doStatement>
          <keyword>do</keyword>
          <identifier>draw</identifier>
          <symbol>(</symbol>
          <expressionList></expressionList>
          <symbol>)</symbol>
          <symbol>;</symbol>
        </doStatement>
        <returnStatement>
          <keyword>return</keyword>
          <expression>
            <term>
              <keyword>this</keyword>
            </term>
          </expression>
          <symbol>;</symbol>
        </returnStatement>
      </statements>
      <symbol>}</symbol>
    </subroutineBody>
  </subroutineDec>
  <subroutineDec>
    <keyword>method</keyword>
    <keyword>void</keyword>
    <identifier>dispose</identifier>
    <symbol>(</symbol>
    <parameterList></parameterList>
    <symbol>)</symbol>
    <subroutineBody>
      <symbol>{</symbol>
      <statements>
        <doStatement>
          <keyword>do</keyword>
          <identifier>Memory</identifier>
          <symbol>.</symbol>
          <identifier>deAlloc</identifier>
          <symbol>(</symbol>
          <expressionList>
            <expression>
              <term>
                <keyword>this</keyword>
              </term>
            </expression>
          </expressionList>
          <symbol>)</symbol>
          <symbol>;</symbol>
        </doStatement>
        <returnStatement>
          <keyword>return</keyword>
          <symbol>;</symbol>
        </returnStatement>
      </statements>
      <symbol>}</symbol>
    </subroutineBody>
  </subroutineDec>
  <subroutineDec>
    <keyword>method</keyword>
    <keyword>void</keyword>
    <identifier>draw</identifier>
    <symbol>(</symbol>
    <parameterList></parameterList>
    <symbol>)</symbol>
    <subroutineBody>
      <symbol>{</symbol>
      <statements>
        <doStatement>
          <keyword>do</keyword>
          <identifier>Screen</identifier>
          <symbol>.</symbol>
          <identifier>setColor</identifier>
          <symbol>(</symbol>
          <expressionList>
            <expression>
              <term>
                <keyword>true</keyword>
              </term>
            </expression>
          </expressionList>
          <symbol>)</symbol>
          <symbol>;</symbol>
        </doStatement>
        <doStatement>
          <keyword>do</keyword>
          <identifier>Screen</identifier>
          <symbol>.</symbol>
          <identifier>drawRectangle</identifier>
          <symbol>(</symbol>
          <expressionList>
            <expression>
              <term>
                <identifier>x</identifier>
              </term>
            </expression>
            <symbol>,</symbol>
            <expression>
              <term>
                <identifier>y</identifier>
              </term>
            </expression>
            <symbol>,</symbol>
            <expression>
              <term>
                <identifier>x</identifier>
              </term>
              <symbol>+</symbol>
              <term>
                <identifier>size</identifier>
              </term>
            </expression>
            <symbol>,</symbol>
            <expression>
              <term>
                <identifier>y</identifier>
              </term>
              <symbol>+</symbol>
              <term>
                <identifier>size</identifier>
              </term>
            </expression>
          </expressionList>
          <symbol>)</symbol>
          <symbol>;</symbol>
        </doStatement>
        <returnStatement>
          <keyword>return</keyword>
          <symbol>;</symbol>
        </returnStatement>
      </statements>
      <symbol>}</symbol>
    </subroutineBody>
  </subroutineDec>
  <subroutineDec>
    <keyword>method</keyword>
    <keyword>void</keyword>
    <identifier>erase</identifier>
    <symbol>(</symbol>
    <parameterList></parameterList>
    <symbol>)</symbol>
    <subroutineBody>
      <symbol>{</symbol>
      <statements>
        <doStatement>
          <keyword>do</keyword>
          <identifier>Screen</identifier>
          <symbol>.</symbol>
          <identifier>setColor</identifier>
          <symbol>(</symbol>
          <expressionList>
            <expression>
              <term>
                <keyword>false</keyword>
              </term>
            </expression>
          </expressionList>
          <symbol>)</symbol>
          <symbol>;</symbol>
        </doStatement>
        <doStatement>
          <keyword>do</keyword>
          <identifier>Screen</identifier>
          <symbol>.</symbol>
          <identifier>drawRectangle</identifier>
          <symbol>(</symbol>
          <expressionList>
            <expression>
              <term>
                <identifier>x</identifier>
              </term>
            </expression>
            <symbol>,</symbol>
            <expression>
              <term>
                <identifier>y</identifier>
              </term>
            </expression>
            <symbol>,</symbol>
            <expression>
              <term>
                <identifier>x</identifier>
              </term>
              <symbol>+</symbol>
              <term>
                <identifier>size</identifier>
              </term>
            </expression>
            <symbol>,</symbol>
            <expression>
              <term>
                <identifier>y</identifier>
              </term>
              <symbol>+</symbol>
              <term>
                <identifier>size</identifier>
              </term>
            </expression>
          </expressionList>
          <symbol>)</symbol>
          <symbol>;</symbol>
        </doStatement>
        <returnStatement>
          <keyword>return</keyword>
          <symbol>;</symbol>
        </returnStatement>
      </statements>
      <symbol>}</symbol>
    </subroutineBody>
  </subroutineDec>
  <subroutineDec>
    <keyword>method</keyword>
    <keyword>void</keyword>
    <identifier>incSize</identifier>
    <symbol>(</symbol>
    <parameterList></parameterList>
    <symbol>)</symbol>
    <subroutineBody>
      <symbol>{</symbol>
      <statements>
        <ifStatement>
          <keyword>if</keyword>
          <symbol>(</symbol>
          <expression>
            <term>
              <symbol>(</symbol>
              <expression>
                <term>
                  <symbol>(</symbol>
                  <expression>
                    <term>
                      <identifier>y</identifier>
                    </term>
                    <symbol>+</symbol>
                    <term>
                      <identifier>size</identifier>
                    </term>
                  </expression>
                  <symbol>)</symbol>
                </term>
                <symbol>&lt;</symbol>
                <term>
                  <integerConstant>254</integerConstant>
                </term>
              </expression>
              <symbol>)</symbol>
            </term>
            <symbol>&amp;</symbol>
            <term>
              <symbol>(</symbol>
              <expression>
                <term>
                  <symbol>(</symbol>
                  <expression>
                    <term>
                      <identifier>x</identifier>
                    </term>
                    <symbol>+</symbol>
                    <term>
                      <identifier>size</identifier>
                    </term>
                  </expression>
                  <symbol>)</symbol>
                </term>
                <symbol>&lt;</symbol>
                <term>
                  <integerConstant>510</integerConstant>
                </term>
              </expression>
              <symbol>)</symbol>
            </term>
          </expression>
          <symbol>)</symbol>
          <symbol>{</symbol>
          <statements>
            <doStatement>
              <keyword>do</keyword>
              <identifier>erase</identifier>
              <symbol>(</symbol>
              <expressionList></expressionList>
              <symbol>)</symbol>
              <symbol>;</symbol>
            </doStatement>
            <letStatement>
              <keyword>let</keyword>
              <identifier>size</identifier>
              <symbol>=</symbol>
              <expression>
                <term>
                  <identifier>size</identifier>
                </term>
                <symbol>+</symbol>
                <term>
                  <integerConstant>2</integerConstant>
                </term>
              </expression>
              <symbol>;</symbol>
            </letStatement>
            <doStatement>
              <keyword>do</keyword>
              <identifier>draw</identifier>
              <symbol>(</symbol>
              <expressionList></expressionList>
              <symbol>)</symbol>
              <symbol>;</symbol>
            </doStatement>
          </statements>
          <symbol>}</symbol>
        </ifStatement>
        <returnStatement>
          <keyword>return</keyword>
          <symbol>;</symbol>
        </returnStatement>
      </statements>
      <symbol>}</symbol>
    </subroutineBody>
  </subroutineDec>
  <subroutineDec>
    <keyword>method</keyword>
    <keyword>void</keyword>
    <identifier>decSize</identifier>
    <symbol>(</symbol>
    <parameterList></parameterList>
    <symbol>)</symbol>
    <subroutineBody>
      <symbol>{</symbol>
      <statements>
        <ifStatement>
          <keyword>if</keyword>
          <symbol>(</symbol>
          <expression>
            <term>
              <identifier>size</identifier>
            </term>
            <symbol>&gt;</symbol>
            <term>
              <integerConstant>2</integerConstant>
            </term>
          </expression>
          <symbol>)</symbol>
          <symbol>{</symbol>
          <statements>
            <doStatement>
              <keyword>do</keyword>
              <identifier>erase</identifier>
              <symbol>(</symbol>
              <expressionList></expressionList>
              <symbol>)</symbol>
              <symbol>;</symbol>
            </doStatement>
            <letStatement>
              <keyword>let</keyword>
              <identifier>size</identifier>
              <symbol>=</symbol>
              <expression>
                <term>
                  <identifier>size</identifier>
                </term>
                <symbol>-</symbol>
                <term>
                  <integerConstant>2</integerConstant>
                </term>
              </expression>
              <symbol>;</symbol>
            </letStatement>
            <doStatement>
              <keyword>do</keyword>
              <identifier>draw</identifier>
              <symbol>(</symbol>
              <expressionList></expressionList>
              <symbol>)</symbol>
              <symbol>;</symbol>
            </doStatement>
          </statements>
          <symbol>}</symbol>
        </ifStatement>
        <returnStatement>
          <keyword>return</keyword>
          <symbol>;</symbol>
        </returnStatement>
      </statements>
      <symbol>}</symbol>
    </subroutineBody>
  </subroutineDec>
  <subroutineDec>
    <keyword>method</keyword>
    <keyword>void</keyword>
    <identifier>moveUp</identifier>
    <symbol>(</symbol>
    <parameterList></parameterList>
    <symbol>)</symbol>
    <subroutineBody>
      <symbol>{</symbol>
      <statements>
        <ifStatement>
          <keyword>if</keyword>
          <symbol>(</symbol>
          <expression>
            <term>
              <identifier>y</identifier>
            </term>
            <symbol>&gt;</symbol>
            <term>
              <integerConstant>1</integerConstant>
            </term>
          </expression>
          <symbol>)</symbol>
          <symbol>{</symbol>
          <statements>
            <doStatement>
              <keyword>do</keyword>
              <identifier>Screen</identifier>
              <symbol>.</symbol>
              <identifier>setColor</identifier>
              <symbol>(</symbol>
              <expressionList>
                <expression>
                  <term>
                    <keyword>false</keyword>
                  </term>
                </expression>
              </expressionList>
              <symbol>)</symbol>
              <symbol>;</symbol>
            </doStatement>
            <doStatement>
              <keyword>do</keyword>
              <identifier>Screen</identifier>
              <symbol>.</symbol>
              <identifier>drawRectangle</identifier>
              <symbol>(</symbol>
              <expressionList>
                <expression>
                  <term>
                    <identifier>x</identifier>
                  </term>
                </expression>
                <symbol>,</symbol>
                <expression>
                  <term>
                    <symbol>(</symbol>
                    <expression>
                      <term>
                        <identifier>y</identifier>
                      </term>
                      <symbol>+</symbol>
                      <term>
                        <identifier>size</identifier>
                      </term>
                    </expression>
                    <symbol>)</symbol>
                  </term>
                  <symbol>-</symbol>
                  <term>
                    <integerConstant>1</integerConstant>
                  </term>
                </expression>
                <symbol>,</symbol>
                <expression>
                  <term>
                    <identifier>x</identifier>
                  </term>
                  <symbol>+</symbol>
                  <term>
                    <identifier>size</identifier>
                  </term>
                </expression>
                <symbol>,</symbol>
                <expression>
                  <term>
                    <identifier>y</identifier>
                  </term>
                  <symbol>+</symbol>
                  <term>
                    <identifier>size</identifier>
                  </term>
                </expression>
              </expressionList>
              <symbol>)</symbol>
              <symbol>;</symbol>
            </doStatement>
            <letStatement>
              <keyword>let</keyword>
              <identifier>y</identifier>
              <symbol>=</symbol>
              <expression>
                <term>
                  <identifier>y</identifier>
                </term>
                <symbol>-</symbol>
                <term>
                  <integerConstant>2</integerConstant>
                </term>
              </expression>
              <symbol>;</symbol>
            </letStatement>
            <doStatement>
              <keyword>do</keyword>
              <identifier>Screen</identifier>
              <symbol>.</symbol>
              <identifier>setColor</identifier>
              <symbol>(</symbol>
              <expressionList>
                <expression>
                  <term>
                    <keyword>true</keyword>
                  </term>
                </expression>
              </expressionList>
              <symbol>)</symbol>
              <symbol>;</symbol>
            </doStatement>
            <doStatement>
              <keyword>do</keyword>
              <identifier>Screen</identifier>
              <symbol>.</symbol>
              <identifier>drawRectangle</identifier>
              <symbol>(</symbol>
              <expressionList>
                <expression>
                  <term>
                    <identifier>x</identifier>
                  </term>
                </expression>
                <symbol>,</symbol>
                <expression>
                  <term>
                    <identifier>y</identifier>
                  </term>
                </expression>
                <symbol>,</symbol>
                <expression>
                  <term>
                    <identifier>x</identifier>
                  </term>
                  <symbol>+</symbol>
                  <term>
                    <identifier>size</identifier>
                  </term>
                </expression>
                <symbol>,</symbol>
                <expression>
                  <term>
                    <identifier>y</identifier>
                  </term>
                  <symbol>+</symbol>
                  <term>
                    <integerConstant>1</integerConstant>
                  </term>
                </expression>
              </expressionList>
              <symbol>)</symbol>
              <symbol>;</symbol>
            </doStatement>
          </statements>
          <symbol>}</symbol>
        </ifStatement>
        <returnStatement>
          <keyword>return</keyword>
          <symbol>;</symbol>
        </returnStatement>
      </statements>
      <symbol>}</symbol>
    </subroutineBody>
  </subroutineDec>
  <subroutineDec>
    <keyword>method</keyword>
    <keyword>void</keyword>
    <identifier>moveDown</identifier>
    <symbol>(</symbol>
    <parameterList></parameterList>
    <symbol>)</symbol>
    <subroutineBody>
      <symbol>{</symbol>
      <statements>
        <ifStatement>
          <keyword>if</keyword>
          <symbol>(</symbol>
          <expression>
            <term>
              <symbol>(</symbol>
              <expression>
                <term>
                  <identifier>y</identifier>
                </term>
                <symbol>+</symbol>
                <term>
                  <identifier>size</identifier>
                </term>
              </expression>
              <symbol>)</symbol>
            </term>
            <symbol>&lt;</symbol>
            <term>
              <integerConstant>254</integerConstant>
            </term>
          </expression>
          <symbol>)</symbol>
          <symbol>{</symbol>
          <statements>
            <doStatement>
              <keyword>do</keyword>
              <identifier>Screen</identifier>
              <symbol>.</symbol>
              <identifier>setColor</identifier>
              <symbol>(</symbol>
              <expressionList>
                <expression>
                  <term>
                    <keyword>false</keyword>
                  </term>
                </expression>
              </expressionList>
              <symbol>)</symbol>
              <symbol>;</symbol>
            </doStatement>
            <doStatement>
              <keyword>do</keyword>
              <identifier>Screen</identifier>
              <symbol>.</symbol>
              <identifier>drawRectangle</identifier>
              <symbol>(</symbol>
              <expressionList>
                <expression>
                  <term>
                    <identifier>x</identifier>
                  </term>
                </expression>
                <symbol>,</symbol>
                <expression>
                  <term>
                    <identifier>y</identifier>
                  </term>
                </expression>
                <symbol>,</symbol>
                <expression>
                  <term>
                    <identifier>x</identifier>
                  </term>
                  <symbol>+</symbol>
                  <term>
                    <identifier>size</identifier>
                  </term>
                </expression>
                <symbol>,</symbol>
                <expression>
                  <term>
                    <identifier>y</identifier>
                  </term>
                  <symbol>+</symbol>
                  <term>
                    <integerConstant>1</integerConstant>
                  </term>
                </expression>
              </expressionList>
              <symbol>)</symbol>
              <symbol>;</symbol>
            </doStatement>
            <letStatement>
              <keyword>let</keyword>
              <identifier>y</identifier>
              <symbol>=</symbol>
              <expression>
                <term>
                  <identifier>y</identifier>
                </term>
                <symbol>+</symbol>
                <term>
                  <integerConstant>2</integerConstant>
                </term>
              </expression>
              <symbol>;</symbol>
            </letStatement>
            <doStatement>
              <keyword>do</keyword>
              <identifier>Screen</identifier>
              <symbol>.</symbol>
              <identifier>setColor</identifier>
              <symbol>(</symbol>
              <expressionList>
                <expression>
                  <term>
                    <keyword>true</keyword>
                  </term>
                </expression>
              </expressionList>
              <symbol>)</symbol>
              <symbol>;</symbol>
            </doStatement>
            <doStatement>
              <keyword>do</keyword>
              <identifier>Screen</identifier>
              <symbol>.</symbol>
              <identifier>drawRectangle</identifier>
              <symbol>(</symbol>
              <expressionList>
                <expression>
                  <term>
                    <identifier>x</identifier>
                  </term>
                </expression>
                <symbol>,</symbol>
                <expression>
                  <term>
                    <symbol>(</symbol>
                    <expression>
                      <term>
                        <identifier>y</identifier>
                      </term>
                      <symbol>+</symbol>
                      <term>
                        <identifier>size</identifier>
                      </term>
                    </expression>
                    <symbol>)</symbol>
                  </term>
                  <symbol>-</symbol>
                  <term>
                    <integerConstant>1</integerConstant>
                  </term>
                </expression>
                <symbol>,</symbol>
                <expression>
                  <term>
                    <identifier>x</identifier>
                  </term>
                  <symbol>+</symbol>
                  <term>
                    <identifier>size</identifier>
                  </term>
                </expression>
                <symbol>,</symbol>
                <expression>
                  <term>
                    <identifier>y</identifier>
                  </term>
                  <symbol>+</symbol>
                  <term>
                    <identifier>size</identifier>
                  </term>
                </expression>
              </expressionList>
              <symbol>)</symbol>
              <symbol>;</symbol>
            </doStatement>
          </statements>
          <symbol>}</symbol>
        </ifStatement>
        <returnStatement>
          <keyword>return</keyword>
          <symbol>;</symbol>
        </returnStatement>
      </statements>
      <symbol>}</symbol>
    </subroutineBody>
  </subroutineDec>
  <subroutineDec>
    <keyword>method</keyword>
    <keyword>void</keyword>
    <identifier>moveLeft</identifier>
    <symbol>(</symbol>
    <parameterList></parameterList>
    <symbol>)</symbol>
    <subroutineBody>
      <symbol>{</symbol>
      <statements>
        <ifStatement>
          <keyword>if</keyword>
          <symbol>(</symbol>
          <expression>
            <term>
              <identifier>x</identifier>
            </term>
            <symbol>&gt;</symbol>
            <term>
              <integerConstant>1</integerConstant>
            </term>
          </expression>
          <symbol>)</symbol>
          <symbol>{</symbol>
          <statements>
            <doStatement>
              <keyword>do</keyword>
              <identifier>Screen</identifier>
              <symbol>.</symbol>
              <identifier>setColor</identifier>
              <symbol>(</symbol>
              <expressionList>
                <expression>
                  <term>
                    <keyword>false</keyword>
                  </term>
                </expression>
              </expressionList>
              <symbol>)</symbol>
              <symbol>;</symbol>
            </doStatement>
            <doStatement>
              <keyword>do</keyword>
              <identifier>Screen</identifier>
              <symbol>.</symbol>
              <identifier>drawRectangle</identifier>
              <symbol>(</symbol>
              <expressionList>
                <expression>
                  <term>
                    <symbol>(</symbol>
                    <expression>
                      <term>
                        <identifier>x</identifier>
                      </term>
                      <symbol>+</symbol>
                      <term>
                        <identifier>size</identifier>
                      </term>
                    </expression>
                    <symbol>)</symbol>
                  </term>
                  <symbol>-</symbol>
                  <term>
                    <integerConstant>1</integerConstant>
                  </term>
                </expression>
                <symbol>,</symbol>
                <expression>
                  <term>
                    <identifier>y</identifier>
                  </term>
                </expression>
                <symbol>,</symbol>
                <expression>
                  <term>
                    <identifier>x</identifier>
                  </term>
                  <symbol>+</symbol>
                  <term>
                    <identifier>size</identifier>
                  </term>
                </expression>
                <symbol>,</symbol>
                <expression>
                  <term>
                    <identifier>y</identifier>
                  </term>
                  <symbol>+</symbol>
                  <term>
                    <identifier>size</identifier>
                  </term>
                </expression>
              </expressionList>
              <symbol>)</symbol>
              <symbol>;</symbol>
            </doStatement>
            <letStatement>
              <keyword>let</keyword>
              <identifier>x</identifier>
              <symbol>=</symbol>
              <expression>
                <term>
                  <identifier>x</identifier>
                </term>
                <symbol>-</symbol>
                <term>
                  <integerConstant>2</integerConstant>
                </term>
              </expression>
              <symbol>;</symbol>
            </letStatement>
            <doStatement>
              <keyword>do</keyword>
              <identifier>Screen</identifier>
              <symbol>.</symbol>
              <identifier>setColor</identifier>
              <symbol>(</symbol>
              <expressionList>
                <expression>
                  <term>
                    <keyword>true</keyword>
                  </term>
                </expression>
              </expressionList>
              <symbol>)</symbol>
              <symbol>;</symbol>
            </doStatement>
            <doStatement>
              <keyword>do</keyword>
              <identifier>Screen</identifier>
              <symbol>.</symbol>
              <identifier>drawRectangle</identifier>
              <symbol>(</symbol>
              <expressionList>
                <expression>
                  <term>
                    <identifier>x</identifier>
                  </term>
                </expression>
                <symbol>,</symbol>
                <expression>
                  <term>
                    <identifier>y</identifier>
                  </term>
                </expression>
                <symbol>,</symbol>
                <expression>
                  <term>
                    <identifier>x</identifier>
                  </term>
                  <symbol>+</symbol>
                  <term>
                    <integerConstant>1</integerConstant>
                  </term>
                </expression>
                <symbol>,</symbol>
                <expression>
                  <term>
                    <identifier>y</identifier>
                  </term>
                  <symbol>+</symbol>
                  <term>
                    <identifier>size</identifier>
                  </term>
                </expression>
              </expressionList>
              <symbol>)</symbol>
              <symbol>;</symbol>
            </doStatement>
          </statements>
          <symbol>}</symbol>
        </ifStatement>
        <returnStatement>
          <keyword>return</keyword>
          <symbol>;</symbol>
        </returnStatement>
      </statements>
      <symbol>}</symbol>
    </subroutineBody>
  </subroutineDec>
  <subroutineDec>
    <keyword>method</keyword>
    <keyword>void</keyword>
    <identifier>moveRight</identifier>
    <symbol>(</symbol>
    <parameterList></parameterList>
    <symbol>)</symbol>
    <subroutineBody>
      <symbol>{</symbol>
      <statements>
        <ifStatement>
          <keyword>if</keyword>
          <symbol>(</symbol>
          <expression>
            <term>
              <symbol>(</symbol>
              <expression>
                <term>
                  <identifier>x</identifier>
                </term>
                <symbol>+</symbol>
                <term>
                  <identifier>size</identifier>
                </term>
              </expression>
              <symbol>)</symbol>
            </term>
            <symbol>&lt;</symbol>
            <term>
              <integerConstant>510</integerConstant>
            </term>
          </expression>
          <symbol>)</symbol>
          <symbol>{</symbol>
          <statements>
            <doStatement>
              <keyword>do</keyword>
              <identifier>Screen</identifier>
              <symbol>.</symbol>
              <identifier>setColor</identifier>
              <symbol>(</symbol>
              <expressionList>
                <expression>
                  <term>
                    <keyword>false</keyword>
                  </term>
                </expression>
              </expressionList>
              <symbol>)</symbol>
              <symbol>;</symbol>
            </doStatement>
            <doStatement>
              <keyword>do</keyword>
              <identifier>Screen</identifier>
              <symbol>.</symbol>
              <identifier>drawRectangle</identifier>
              <symbol>(</symbol>
              <expressionList>
                <expression>
                  <term>
                    <identifier>x</identifier>
                  </term>
                </expression>
                <symbol>,</symbol>
                <expression>
                  <term>
                    <identifier>y</identifier>
                  </term>
                </expression>
                <symbol>,</symbol>
                <expression>
                  <term>
                    <identifier>x</identifier>
                  </term>
                  <symbol>+</symbol>
                  <term>
                    <integerConstant>1</integerConstant>
                  </term>
                </expression>
                <symbol>,</symbol>
                <expression>
                  <term>
                    <identifier>y</identifier>
                  </term>
                  <symbol>+</symbol>
                  <term>
                    <identifier>size</identifier>
                  </term>
                </expression>
              </expressionList>
              <symbol>)</symbol>
              <symbol>;</symbol>
            </doStatement>
            <letStatement>
              <keyword>let</keyword>
              <identifier>x</identifier>
              <symbol>=</symbol>
              <expression>
                <term>
                  <identifier>x</identifier>
                </term>
                <symbol>+</symbol>
                <term>
                  <integerConstant>2</integerConstant>
                </term>
              </expression>
              <symbol>;</symbol>
            </letStatement>
            <doStatement>
              <keyword>do</keyword>
              <identifier>Screen</identifier>
              <symbol>.</symbol>
              <identifier>setColor</identifier>
              <symbol>(</symbol>
              <expressionList>
                <expression>
                  <term>
                    <keyword>true</keyword>
                  </term>
                </expression>
              </expressionList>
              <symbol>)</symbol>
              <symbol>;</symbol>
            </doStatement>
            <doStatement>
              <keyword>do</keyword>
              <identifier>Screen</identifier>
              <symbol>.</symbol>
              <identifier>drawRectangle</identifier>
              <symbol>(</symbol>
              <expressionList>
                <expression>
                  <term>
                    <symbol>(</symbol>
                    <expression>
                      <term>
                        <identifier>x</identifier>
                      </term>
                      <symbol>+</symbol>
                      <term>
                        <identifier>size</identifier>
                      </term>
                    </expression>
                    <symbol>)</symbol>
                  </term>
                  <symbol>-</symbol>
                  <term>
                    <integerConstant>1</integerConstant>
                  </term>
                </expression>
                <symbol>,</symbol>
                <expression>
                  <term>
                    <identifier>y</identifier>
                  </term>
                </expression>
                <symbol>,</symbol>
                <expression>
                  <term>
                    <identifier>x</identifier>
                  </term>
                  <symbol>+</symbol>
                  <term>
                    <identifier>size</identifier>
                  </term>
                </expression>
                <symbol>,</symbol>
                <expression>
                  <term>
                    <identifier>y</identifier>
                  </term>
                  <symbol>+</symbol>
                  <term>
                    <identifier>size</identifier>
                  </term>
                </expression>
              </expressionList>
              <symbol>)</symbol>
              <symbol>;</symbol>
            </doStatement>
          </statements>
          <symbol>}</symbol>
        </ifStatement>
        <returnStatement>
          <keyword>return</keyword>
          <symbol>;</symbol>
        </returnStatement>
      </statements>
      <symbol>}</symbol>
    </subroutineBody>
  </subroutineDec>
  <symbol>}</symbol>
</class>`

			reader := strings.NewReader(inputTokensXml)
			testTokens := tokens.LoadTokensFromXML(reader)

			var buf bytes.Buffer
			xmlEnc := xml.NewEncoder(&buf)

			// Act
			ce, err := New(testTokens, xmlEnc)
			assert.NoError(t, err)
			err = ce.CompileClass()

			// Assert
			assert.NoError(t, err)

			assert.Equal(t, expectedXml, buf.String())
		})

		t.Run("should return XML when compiling SquareGame.jack", func(t *testing.T) {

			// Arrange
			inputTokensXml := `<tokens>
  <keyword> class </keyword>
  <identifier> SquareGame </identifier>
  <symbol> { </symbol>
  <keyword> field </keyword>
  <identifier> Square </identifier>
  <identifier> square </identifier>
  <symbol> ; </symbol>
  <keyword> field </keyword>
  <keyword> int </keyword>
  <identifier> direction </identifier>
  <symbol> ; </symbol>
  <keyword> constructor </keyword>
  <identifier> SquareGame </identifier>
  <identifier> new </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> let </keyword>
  <identifier> square </identifier>
  <symbol> = </symbol>
  <identifier> Square </identifier>
  <symbol> . </symbol>
  <identifier> new </identifier>
  <symbol> ( </symbol>
  <integerConstant> 0 </integerConstant>
  <symbol> , </symbol>
  <integerConstant> 0 </integerConstant>
  <symbol> , </symbol>
  <integerConstant> 30 </integerConstant>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> let </keyword>
  <identifier> direction </identifier>
  <symbol> = </symbol>
  <integerConstant> 0 </integerConstant>
  <symbol> ; </symbol>
  <keyword> return </keyword>
  <keyword> this </keyword>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> method </keyword>
  <keyword> void </keyword>
  <identifier> dispose </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> do </keyword>
  <identifier> square </identifier>
  <symbol> . </symbol>
  <identifier> dispose </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> do </keyword>
  <identifier> Memory </identifier>
  <symbol> . </symbol>
  <identifier> deAlloc </identifier>
  <symbol> ( </symbol>
  <keyword> this </keyword>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> return </keyword>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> method </keyword>
  <keyword> void </keyword>
  <identifier> moveSquare </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> if </keyword>
  <symbol> ( </symbol>
  <identifier> direction </identifier>
  <symbol> = </symbol>
  <integerConstant> 1 </integerConstant>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> do </keyword>
  <identifier> square </identifier>
  <symbol> . </symbol>
  <identifier> moveUp </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> if </keyword>
  <symbol> ( </symbol>
  <identifier> direction </identifier>
  <symbol> = </symbol>
  <integerConstant> 2 </integerConstant>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> do </keyword>
  <identifier> square </identifier>
  <symbol> . </symbol>
  <identifier> moveDown </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> if </keyword>
  <symbol> ( </symbol>
  <identifier> direction </identifier>
  <symbol> = </symbol>
  <integerConstant> 3 </integerConstant>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> do </keyword>
  <identifier> square </identifier>
  <symbol> . </symbol>
  <identifier> moveLeft </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> if </keyword>
  <symbol> ( </symbol>
  <identifier> direction </identifier>
  <symbol> = </symbol>
  <integerConstant> 4 </integerConstant>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> do </keyword>
  <identifier> square </identifier>
  <symbol> . </symbol>
  <identifier> moveRight </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> do </keyword>
  <identifier> Sys </identifier>
  <symbol> . </symbol>
  <identifier> wait </identifier>
  <symbol> ( </symbol>
  <integerConstant> 5 </integerConstant>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> return </keyword>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> method </keyword>
  <keyword> void </keyword>
  <identifier> run </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> var </keyword>
  <keyword> char </keyword>
  <identifier> key </identifier>
  <symbol> ; </symbol>
  <keyword> var </keyword>
  <keyword> boolean </keyword>
  <identifier> exit </identifier>
  <symbol> ; </symbol>
  <keyword> let </keyword>
  <identifier> exit </identifier>
  <symbol> = </symbol>
  <keyword> false </keyword>
  <symbol> ; </symbol>
  <keyword> while </keyword>
  <symbol> ( </symbol>
  <symbol> ~ </symbol>
  <identifier> exit </identifier>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> while </keyword>
  <symbol> ( </symbol>
  <identifier> key </identifier>
  <symbol> = </symbol>
  <integerConstant> 0 </integerConstant>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> let </keyword>
  <identifier> key </identifier>
  <symbol> = </symbol>
  <identifier> Keyboard </identifier>
  <symbol> . </symbol>
  <identifier> keyPressed </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> do </keyword>
  <identifier> moveSquare </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> if </keyword>
  <symbol> ( </symbol>
  <identifier> key </identifier>
  <symbol> = </symbol>
  <integerConstant> 81 </integerConstant>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> let </keyword>
  <identifier> exit </identifier>
  <symbol> = </symbol>
  <keyword> true </keyword>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> if </keyword>
  <symbol> ( </symbol>
  <identifier> key </identifier>
  <symbol> = </symbol>
  <integerConstant> 90 </integerConstant>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> do </keyword>
  <identifier> square </identifier>
  <symbol> . </symbol>
  <identifier> decSize </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> if </keyword>
  <symbol> ( </symbol>
  <identifier> key </identifier>
  <symbol> = </symbol>
  <integerConstant> 88 </integerConstant>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> do </keyword>
  <identifier> square </identifier>
  <symbol> . </symbol>
  <identifier> incSize </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> if </keyword>
  <symbol> ( </symbol>
  <identifier> key </identifier>
  <symbol> = </symbol>
  <integerConstant> 131 </integerConstant>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> let </keyword>
  <identifier> direction </identifier>
  <symbol> = </symbol>
  <integerConstant> 1 </integerConstant>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> if </keyword>
  <symbol> ( </symbol>
  <identifier> key </identifier>
  <symbol> = </symbol>
  <integerConstant> 133 </integerConstant>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> let </keyword>
  <identifier> direction </identifier>
  <symbol> = </symbol>
  <integerConstant> 2 </integerConstant>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> if </keyword>
  <symbol> ( </symbol>
  <identifier> key </identifier>
  <symbol> = </symbol>
  <integerConstant> 130 </integerConstant>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> let </keyword>
  <identifier> direction </identifier>
  <symbol> = </symbol>
  <integerConstant> 3 </integerConstant>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> if </keyword>
  <symbol> ( </symbol>
  <identifier> key </identifier>
  <symbol> = </symbol>
  <integerConstant> 132 </integerConstant>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> let </keyword>
  <identifier> direction </identifier>
  <symbol> = </symbol>
  <integerConstant> 4 </integerConstant>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> while </keyword>
  <symbol> ( </symbol>
  <symbol> ~ </symbol>
  <symbol> ( </symbol>
  <identifier> key </identifier>
  <symbol> = </symbol>
  <integerConstant> 0 </integerConstant>
  <symbol> ) </symbol>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> let </keyword>
  <identifier> key </identifier>
  <symbol> = </symbol>
  <identifier> Keyboard </identifier>
  <symbol> . </symbol>
  <identifier> keyPressed </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> do </keyword>
  <identifier> moveSquare </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <symbol> } </symbol>
  <keyword> return </keyword>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <symbol> } </symbol>
</tokens>
`

			expectedXml := `<class>
  <keyword>class</keyword>
  <identifier>SquareGame</identifier>
  <symbol>{</symbol>
  <classVarDec>
    <keyword>field</keyword>
    <identifier>Square</identifier>
    <identifier>square</identifier>
    <symbol>;</symbol>
  </classVarDec>
  <classVarDec>
    <keyword>field</keyword>
    <keyword>int</keyword>
    <identifier>direction</identifier>
    <symbol>;</symbol>
  </classVarDec>
  <subroutineDec>
    <keyword>constructor</keyword>
    <identifier>SquareGame</identifier>
    <identifier>new</identifier>
    <symbol>(</symbol>
    <parameterList></parameterList>
    <symbol>)</symbol>
    <subroutineBody>
      <symbol>{</symbol>
      <statements>
        <letStatement>
          <keyword>let</keyword>
          <identifier>square</identifier>
          <symbol>=</symbol>
          <expression>
            <term>
              <identifier>Square</identifier>
              <symbol>.</symbol>
              <identifier>new</identifier>
              <symbol>(</symbol>
              <expressionList>
                <expression>
                  <term>
                    <integerConstant>0</integerConstant>
                  </term>
                </expression>
                <symbol>,</symbol>
                <expression>
                  <term>
                    <integerConstant>0</integerConstant>
                  </term>
                </expression>
                <symbol>,</symbol>
                <expression>
                  <term>
                    <integerConstant>30</integerConstant>
                  </term>
                </expression>
              </expressionList>
              <symbol>)</symbol>
            </term>
          </expression>
          <symbol>;</symbol>
        </letStatement>
        <letStatement>
          <keyword>let</keyword>
          <identifier>direction</identifier>
          <symbol>=</symbol>
          <expression>
            <term>
              <integerConstant>0</integerConstant>
            </term>
          </expression>
          <symbol>;</symbol>
        </letStatement>
        <returnStatement>
          <keyword>return</keyword>
          <expression>
            <term>
              <keyword>this</keyword>
            </term>
          </expression>
          <symbol>;</symbol>
        </returnStatement>
      </statements>
      <symbol>}</symbol>
    </subroutineBody>
  </subroutineDec>
  <subroutineDec>
    <keyword>method</keyword>
    <keyword>void</keyword>
    <identifier>dispose</identifier>
    <symbol>(</symbol>
    <parameterList></parameterList>
    <symbol>)</symbol>
    <subroutineBody>
      <symbol>{</symbol>
      <statements>
        <doStatement>
          <keyword>do</keyword>
          <identifier>square</identifier>
          <symbol>.</symbol>
          <identifier>dispose</identifier>
          <symbol>(</symbol>
          <expressionList></expressionList>
          <symbol>)</symbol>
          <symbol>;</symbol>
        </doStatement>
        <doStatement>
          <keyword>do</keyword>
          <identifier>Memory</identifier>
          <symbol>.</symbol>
          <identifier>deAlloc</identifier>
          <symbol>(</symbol>
          <expressionList>
            <expression>
              <term>
                <keyword>this</keyword>
              </term>
            </expression>
          </expressionList>
          <symbol>)</symbol>
          <symbol>;</symbol>
        </doStatement>
        <returnStatement>
          <keyword>return</keyword>
          <symbol>;</symbol>
        </returnStatement>
      </statements>
      <symbol>}</symbol>
    </subroutineBody>
  </subroutineDec>
  <subroutineDec>
    <keyword>method</keyword>
    <keyword>void</keyword>
    <identifier>moveSquare</identifier>
    <symbol>(</symbol>
    <parameterList></parameterList>
    <symbol>)</symbol>
    <subroutineBody>
      <symbol>{</symbol>
      <statements>
        <ifStatement>
          <keyword>if</keyword>
          <symbol>(</symbol>
          <expression>
            <term>
              <identifier>direction</identifier>
            </term>
            <symbol>=</symbol>
            <term>
              <integerConstant>1</integerConstant>
            </term>
          </expression>
          <symbol>)</symbol>
          <symbol>{</symbol>
          <statements>
            <doStatement>
              <keyword>do</keyword>
              <identifier>square</identifier>
              <symbol>.</symbol>
              <identifier>moveUp</identifier>
              <symbol>(</symbol>
              <expressionList></expressionList>
              <symbol>)</symbol>
              <symbol>;</symbol>
            </doStatement>
          </statements>
          <symbol>}</symbol>
        </ifStatement>
        <ifStatement>
          <keyword>if</keyword>
          <symbol>(</symbol>
          <expression>
            <term>
              <identifier>direction</identifier>
            </term>
            <symbol>=</symbol>
            <term>
              <integerConstant>2</integerConstant>
            </term>
          </expression>
          <symbol>)</symbol>
          <symbol>{</symbol>
          <statements>
            <doStatement>
              <keyword>do</keyword>
              <identifier>square</identifier>
              <symbol>.</symbol>
              <identifier>moveDown</identifier>
              <symbol>(</symbol>
              <expressionList></expressionList>
              <symbol>)</symbol>
              <symbol>;</symbol>
            </doStatement>
          </statements>
          <symbol>}</symbol>
        </ifStatement>
        <ifStatement>
          <keyword>if</keyword>
          <symbol>(</symbol>
          <expression>
            <term>
              <identifier>direction</identifier>
            </term>
            <symbol>=</symbol>
            <term>
              <integerConstant>3</integerConstant>
            </term>
          </expression>
          <symbol>)</symbol>
          <symbol>{</symbol>
          <statements>
            <doStatement>
              <keyword>do</keyword>
              <identifier>square</identifier>
              <symbol>.</symbol>
              <identifier>moveLeft</identifier>
              <symbol>(</symbol>
              <expressionList></expressionList>
              <symbol>)</symbol>
              <symbol>;</symbol>
            </doStatement>
          </statements>
          <symbol>}</symbol>
        </ifStatement>
        <ifStatement>
          <keyword>if</keyword>
          <symbol>(</symbol>
          <expression>
            <term>
              <identifier>direction</identifier>
            </term>
            <symbol>=</symbol>
            <term>
              <integerConstant>4</integerConstant>
            </term>
          </expression>
          <symbol>)</symbol>
          <symbol>{</symbol>
          <statements>
            <doStatement>
              <keyword>do</keyword>
              <identifier>square</identifier>
              <symbol>.</symbol>
              <identifier>moveRight</identifier>
              <symbol>(</symbol>
              <expressionList></expressionList>
              <symbol>)</symbol>
              <symbol>;</symbol>
            </doStatement>
          </statements>
          <symbol>}</symbol>
        </ifStatement>
        <doStatement>
          <keyword>do</keyword>
          <identifier>Sys</identifier>
          <symbol>.</symbol>
          <identifier>wait</identifier>
          <symbol>(</symbol>
          <expressionList>
            <expression>
              <term>
                <integerConstant>5</integerConstant>
              </term>
            </expression>
          </expressionList>
          <symbol>)</symbol>
          <symbol>;</symbol>
        </doStatement>
        <returnStatement>
          <keyword>return</keyword>
          <symbol>;</symbol>
        </returnStatement>
      </statements>
      <symbol>}</symbol>
    </subroutineBody>
  </subroutineDec>
  <subroutineDec>
    <keyword>method</keyword>
    <keyword>void</keyword>
    <identifier>run</identifier>
    <symbol>(</symbol>
    <parameterList></parameterList>
    <symbol>)</symbol>
    <subroutineBody>
      <symbol>{</symbol>
      <varDec>
        <keyword>var</keyword>
        <keyword>char</keyword>
        <identifier>key</identifier>
        <symbol>;</symbol>
      </varDec>
      <varDec>
        <keyword>var</keyword>
        <keyword>boolean</keyword>
        <identifier>exit</identifier>
        <symbol>;</symbol>
      </varDec>
      <statements>
        <letStatement>
          <keyword>let</keyword>
          <identifier>exit</identifier>
          <symbol>=</symbol>
          <expression>
            <term>
              <keyword>false</keyword>
            </term>
          </expression>
          <symbol>;</symbol>
        </letStatement>
        <whileStatement>
          <keyword>while</keyword>
          <symbol>(</symbol>
          <expression>
            <term>
              <symbol>~</symbol>
              <term>
                <identifier>exit</identifier>
              </term>
            </term>
          </expression>
          <symbol>)</symbol>
          <symbol>{</symbol>
          <statements>
            <whileStatement>
              <keyword>while</keyword>
              <symbol>(</symbol>
              <expression>
                <term>
                  <identifier>key</identifier>
                </term>
                <symbol>=</symbol>
                <term>
                  <integerConstant>0</integerConstant>
                </term>
              </expression>
              <symbol>)</symbol>
              <symbol>{</symbol>
              <statements>
                <letStatement>
                  <keyword>let</keyword>
                  <identifier>key</identifier>
                  <symbol>=</symbol>
                  <expression>
                    <term>
                      <identifier>Keyboard</identifier>
                      <symbol>.</symbol>
                      <identifier>keyPressed</identifier>
                      <symbol>(</symbol>
                      <expressionList></expressionList>
                      <symbol>)</symbol>
                    </term>
                  </expression>
                  <symbol>;</symbol>
                </letStatement>
                <doStatement>
                  <keyword>do</keyword>
                  <identifier>moveSquare</identifier>
                  <symbol>(</symbol>
                  <expressionList></expressionList>
                  <symbol>)</symbol>
                  <symbol>;</symbol>
                </doStatement>
              </statements>
              <symbol>}</symbol>
            </whileStatement>
            <ifStatement>
              <keyword>if</keyword>
              <symbol>(</symbol>
              <expression>
                <term>
                  <identifier>key</identifier>
                </term>
                <symbol>=</symbol>
                <term>
                  <integerConstant>81</integerConstant>
                </term>
              </expression>
              <symbol>)</symbol>
              <symbol>{</symbol>
              <statements>
                <letStatement>
                  <keyword>let</keyword>
                  <identifier>exit</identifier>
                  <symbol>=</symbol>
                  <expression>
                    <term>
                      <keyword>true</keyword>
                    </term>
                  </expression>
                  <symbol>;</symbol>
                </letStatement>
              </statements>
              <symbol>}</symbol>
            </ifStatement>
            <ifStatement>
              <keyword>if</keyword>
              <symbol>(</symbol>
              <expression>
                <term>
                  <identifier>key</identifier>
                </term>
                <symbol>=</symbol>
                <term>
                  <integerConstant>90</integerConstant>
                </term>
              </expression>
              <symbol>)</symbol>
              <symbol>{</symbol>
              <statements>
                <doStatement>
                  <keyword>do</keyword>
                  <identifier>square</identifier>
                  <symbol>.</symbol>
                  <identifier>decSize</identifier>
                  <symbol>(</symbol>
                  <expressionList></expressionList>
                  <symbol>)</symbol>
                  <symbol>;</symbol>
                </doStatement>
              </statements>
              <symbol>}</symbol>
            </ifStatement>
            <ifStatement>
              <keyword>if</keyword>
              <symbol>(</symbol>
              <expression>
                <term>
                  <identifier>key</identifier>
                </term>
                <symbol>=</symbol>
                <term>
                  <integerConstant>88</integerConstant>
                </term>
              </expression>
              <symbol>)</symbol>
              <symbol>{</symbol>
              <statements>
                <doStatement>
                  <keyword>do</keyword>
                  <identifier>square</identifier>
                  <symbol>.</symbol>
                  <identifier>incSize</identifier>
                  <symbol>(</symbol>
                  <expressionList></expressionList>
                  <symbol>)</symbol>
                  <symbol>;</symbol>
                </doStatement>
              </statements>
              <symbol>}</symbol>
            </ifStatement>
            <ifStatement>
              <keyword>if</keyword>
              <symbol>(</symbol>
              <expression>
                <term>
                  <identifier>key</identifier>
                </term>
                <symbol>=</symbol>
                <term>
                  <integerConstant>131</integerConstant>
                </term>
              </expression>
              <symbol>)</symbol>
              <symbol>{</symbol>
              <statements>
                <letStatement>
                  <keyword>let</keyword>
                  <identifier>direction</identifier>
                  <symbol>=</symbol>
                  <expression>
                    <term>
                      <integerConstant>1</integerConstant>
                    </term>
                  </expression>
                  <symbol>;</symbol>
                </letStatement>
              </statements>
              <symbol>}</symbol>
            </ifStatement>
            <ifStatement>
              <keyword>if</keyword>
              <symbol>(</symbol>
              <expression>
                <term>
                  <identifier>key</identifier>
                </term>
                <symbol>=</symbol>
                <term>
                  <integerConstant>133</integerConstant>
                </term>
              </expression>
              <symbol>)</symbol>
              <symbol>{</symbol>
              <statements>
                <letStatement>
                  <keyword>let</keyword>
                  <identifier>direction</identifier>
                  <symbol>=</symbol>
                  <expression>
                    <term>
                      <integerConstant>2</integerConstant>
                    </term>
                  </expression>
                  <symbol>;</symbol>
                </letStatement>
              </statements>
              <symbol>}</symbol>
            </ifStatement>
            <ifStatement>
              <keyword>if</keyword>
              <symbol>(</symbol>
              <expression>
                <term>
                  <identifier>key</identifier>
                </term>
                <symbol>=</symbol>
                <term>
                  <integerConstant>130</integerConstant>
                </term>
              </expression>
              <symbol>)</symbol>
              <symbol>{</symbol>
              <statements>
                <letStatement>
                  <keyword>let</keyword>
                  <identifier>direction</identifier>
                  <symbol>=</symbol>
                  <expression>
                    <term>
                      <integerConstant>3</integerConstant>
                    </term>
                  </expression>
                  <symbol>;</symbol>
                </letStatement>
              </statements>
              <symbol>}</symbol>
            </ifStatement>
            <ifStatement>
              <keyword>if</keyword>
              <symbol>(</symbol>
              <expression>
                <term>
                  <identifier>key</identifier>
                </term>
                <symbol>=</symbol>
                <term>
                  <integerConstant>132</integerConstant>
                </term>
              </expression>
              <symbol>)</symbol>
              <symbol>{</symbol>
              <statements>
                <letStatement>
                  <keyword>let</keyword>
                  <identifier>direction</identifier>
                  <symbol>=</symbol>
                  <expression>
                    <term>
                      <integerConstant>4</integerConstant>
                    </term>
                  </expression>
                  <symbol>;</symbol>
                </letStatement>
              </statements>
              <symbol>}</symbol>
            </ifStatement>
            <whileStatement>
              <keyword>while</keyword>
              <symbol>(</symbol>
              <expression>
                <term>
                  <symbol>~</symbol>
                  <term>
                    <symbol>(</symbol>
                    <expression>
                      <term>
                        <identifier>key</identifier>
                      </term>
                      <symbol>=</symbol>
                      <term>
                        <integerConstant>0</integerConstant>
                      </term>
                    </expression>
                    <symbol>)</symbol>
                  </term>
                </term>
              </expression>
              <symbol>)</symbol>
              <symbol>{</symbol>
              <statements>
                <letStatement>
                  <keyword>let</keyword>
                  <identifier>key</identifier>
                  <symbol>=</symbol>
                  <expression>
                    <term>
                      <identifier>Keyboard</identifier>
                      <symbol>.</symbol>
                      <identifier>keyPressed</identifier>
                      <symbol>(</symbol>
                      <expressionList></expressionList>
                      <symbol>)</symbol>
                    </term>
                  </expression>
                  <symbol>;</symbol>
                </letStatement>
                <doStatement>
                  <keyword>do</keyword>
                  <identifier>moveSquare</identifier>
                  <symbol>(</symbol>
                  <expressionList></expressionList>
                  <symbol>)</symbol>
                  <symbol>;</symbol>
                </doStatement>
              </statements>
              <symbol>}</symbol>
            </whileStatement>
          </statements>
          <symbol>}</symbol>
        </whileStatement>
        <returnStatement>
          <keyword>return</keyword>
          <symbol>;</symbol>
        </returnStatement>
      </statements>
      <symbol>}</symbol>
    </subroutineBody>
  </subroutineDec>
  <symbol>}</symbol>
</class>`

			reader := strings.NewReader(inputTokensXml)
			testTokens := tokens.LoadTokensFromXML(reader)

			var buf bytes.Buffer
			xmlEnc := xml.NewEncoder(&buf)

			// Act
			ce, err := New(testTokens, xmlEnc)
			assert.NoError(t, err)
			err = ce.CompileClass()

			// Assert
			assert.NoError(t, err)

			assert.Equal(t, expectedXml, buf.String())
		})
	})

	t.Run("should return XML when compiling ExpressionLessSquare", func(t *testing.T) {
		t.Run("should return XML when compiling Main.jack", func(t *testing.T) {

			// Arrange
			inputTokensXml := `<tokens>
  <keyword> class </keyword>
  <identifier> Main </identifier>
  <symbol> { </symbol>
  <keyword> static </keyword>
  <keyword> boolean </keyword>
  <identifier> test </identifier>
  <symbol> ; </symbol>
  <keyword> function </keyword>
  <keyword> void </keyword>
  <identifier> main </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> var </keyword>
  <identifier> SquareGame </identifier>
  <identifier> game </identifier>
  <symbol> ; </symbol>
  <keyword> let </keyword>
  <identifier> game </identifier>
  <symbol> = </symbol>
  <identifier> SquareGame </identifier>
  <symbol> . </symbol>
  <identifier> new </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> do </keyword>
  <identifier> game </identifier>
  <symbol> . </symbol>
  <identifier> run </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> do </keyword>
  <identifier> game </identifier>
  <symbol> . </symbol>
  <identifier> dispose </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> return </keyword>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> function </keyword>
  <keyword> void </keyword>
  <identifier> more </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> var </keyword>
  <keyword> int </keyword>
  <identifier> i </identifier>
  <symbol> , </symbol>
  <identifier> j </identifier>
  <symbol> ; </symbol>
  <keyword> var </keyword>
  <identifier> String </identifier>
  <identifier> s </identifier>
  <symbol> ; </symbol>
  <keyword> var </keyword>
  <identifier> Array </identifier>
  <identifier> a </identifier>
  <symbol> ; </symbol>
  <keyword> if </keyword>
  <symbol> ( </symbol>
  <keyword> false </keyword>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> let </keyword>
  <identifier> s </identifier>
  <symbol> = </symbol>
  <stringConstant> string constant </stringConstant>
  <symbol> ; </symbol>
  <keyword> let </keyword>
  <identifier> s </identifier>
  <symbol> = </symbol>
  <keyword> null </keyword>
  <symbol> ; </symbol>
  <keyword> let </keyword>
  <identifier> a </identifier>
  <symbol> [ </symbol>
  <integerConstant> 1 </integerConstant>
  <symbol> ] </symbol>
  <symbol> = </symbol>
  <identifier> a </identifier>
  <symbol> [ </symbol>
  <integerConstant> 2 </integerConstant>
  <symbol> ] </symbol>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> else </keyword>
  <symbol> { </symbol>
  <keyword> let </keyword>
  <identifier> i </identifier>
  <symbol> = </symbol>
  <identifier> i </identifier>
  <symbol> * </symbol>
  <symbol> ( </symbol>
  <symbol> - </symbol>
  <identifier> j </identifier>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> let </keyword>
  <identifier> j </identifier>
  <symbol> = </symbol>
  <identifier> j </identifier>
  <symbol> / </symbol>
  <symbol> ( </symbol>
  <symbol> - </symbol>
  <integerConstant> 2 </integerConstant>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> let </keyword>
  <identifier> i </identifier>
  <symbol> = </symbol>
  <identifier> i </identifier>
  <symbol> | </symbol>
  <identifier> j </identifier>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> return </keyword>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <symbol> } </symbol>
</tokens>
`

			expectedXml := `<class>
  <keyword>class</keyword>
  <identifier>Main</identifier>
  <symbol>{</symbol>
  <classVarDec>
    <keyword>static</keyword>
    <keyword>boolean</keyword>
    <identifier>test</identifier>
    <symbol>;</symbol>
  </classVarDec>
  <subroutineDec>
    <keyword>function</keyword>
    <keyword>void</keyword>
    <identifier>main</identifier>
    <symbol>(</symbol>
    <parameterList></parameterList>
    <symbol>)</symbol>
    <subroutineBody>
      <symbol>{</symbol>
      <varDec>
        <keyword>var</keyword>
        <identifier>SquareGame</identifier>
        <identifier>game</identifier>
        <symbol>;</symbol>
      </varDec>
      <statements>
        <letStatement>
          <keyword>let</keyword>
          <identifier>game</identifier>
          <symbol>=</symbol>
          <expression>
            <term>
              <identifier>SquareGame</identifier>
              <symbol>.</symbol>
              <identifier>new</identifier>
              <symbol>(</symbol>
              <expressionList></expressionList>
              <symbol>)</symbol>
            </term>
          </expression>
          <symbol>;</symbol>
        </letStatement>
        <doStatement>
          <keyword>do</keyword>
          <identifier>game</identifier>
          <symbol>.</symbol>
          <identifier>run</identifier>
          <symbol>(</symbol>
          <expressionList></expressionList>
          <symbol>)</symbol>
          <symbol>;</symbol>
        </doStatement>
        <doStatement>
          <keyword>do</keyword>
          <identifier>game</identifier>
          <symbol>.</symbol>
          <identifier>dispose</identifier>
          <symbol>(</symbol>
          <expressionList></expressionList>
          <symbol>)</symbol>
          <symbol>;</symbol>
        </doStatement>
        <returnStatement>
          <keyword>return</keyword>
          <symbol>;</symbol>
        </returnStatement>
      </statements>
      <symbol>}</symbol>
    </subroutineBody>
  </subroutineDec>
  <subroutineDec>
    <keyword>function</keyword>
    <keyword>void</keyword>
    <identifier>more</identifier>
    <symbol>(</symbol>
    <parameterList></parameterList>
    <symbol>)</symbol>
    <subroutineBody>
      <symbol>{</symbol>
      <varDec>
        <keyword>var</keyword>
        <keyword>int</keyword>
        <identifier>i</identifier>
        <symbol>,</symbol>
        <identifier>j</identifier>
        <symbol>;</symbol>
      </varDec>
      <varDec>
        <keyword>var</keyword>
        <identifier>String</identifier>
        <identifier>s</identifier>
        <symbol>;</symbol>
      </varDec>
      <varDec>
        <keyword>var</keyword>
        <identifier>Array</identifier>
        <identifier>a</identifier>
        <symbol>;</symbol>
      </varDec>
      <statements>
        <ifStatement>
          <keyword>if</keyword>
          <symbol>(</symbol>
          <expression>
            <term>
              <keyword>false</keyword>
            </term>
          </expression>
          <symbol>)</symbol>
          <symbol>{</symbol>
          <statements>
            <letStatement>
              <keyword>let</keyword>
              <identifier>s</identifier>
              <symbol>=</symbol>
              <expression>
                <term>
                  <stringConstant>string constant</stringConstant>
                </term>
              </expression>
              <symbol>;</symbol>
            </letStatement>
            <letStatement>
              <keyword>let</keyword>
              <identifier>s</identifier>
              <symbol>=</symbol>
              <expression>
                <term>
                  <keyword>null</keyword>
                </term>
              </expression>
              <symbol>;</symbol>
            </letStatement>
            <letStatement>
              <keyword>let</keyword>
              <identifier>a</identifier>
              <symbol>[</symbol>
              <expression>
                <term>
                  <integerConstant>1</integerConstant>
                </term>
              </expression>
              <symbol>]</symbol>
              <symbol>=</symbol>
              <expression>
                <term>
                  <identifier>a</identifier>
                  <symbol>[</symbol>
                  <expression>
                    <term>
                      <integerConstant>2</integerConstant>
                    </term>
                  </expression>
                  <symbol>]</symbol>
                </term>
              </expression>
              <symbol>;</symbol>
            </letStatement>
          </statements>
          <symbol>}</symbol>
          <keyword>else</keyword>
          <symbol>{</symbol>
          <statements>
            <letStatement>
              <keyword>let</keyword>
              <identifier>i</identifier>
              <symbol>=</symbol>
              <expression>
                <term>
                  <identifier>i</identifier>
                </term>
                <symbol>*</symbol>
                <term>
                  <symbol>(</symbol>
                  <expression>
                    <term>
                      <symbol>-</symbol>
                      <term>
                        <identifier>j</identifier>
                      </term>
                    </term>
                  </expression>
                  <symbol>)</symbol>
                </term>
              </expression>
              <symbol>;</symbol>
            </letStatement>
            <letStatement>
              <keyword>let</keyword>
              <identifier>j</identifier>
              <symbol>=</symbol>
              <expression>
                <term>
                  <identifier>j</identifier>
                </term>
                <symbol>/</symbol>
                <term>
                  <symbol>(</symbol>
                  <expression>
                    <term>
                      <symbol>-</symbol>
                      <term>
                        <integerConstant>2</integerConstant>
                      </term>
                    </term>
                  </expression>
                  <symbol>)</symbol>
                </term>
              </expression>
              <symbol>;</symbol>
            </letStatement>
            <letStatement>
              <keyword>let</keyword>
              <identifier>i</identifier>
              <symbol>=</symbol>
              <expression>
                <term>
                  <identifier>i</identifier>
                </term>
                <symbol>|</symbol>
                <term>
                  <identifier>j</identifier>
                </term>
              </expression>
              <symbol>;</symbol>
            </letStatement>
          </statements>
          <symbol>}</symbol>
        </ifStatement>
        <returnStatement>
          <keyword>return</keyword>
          <symbol>;</symbol>
        </returnStatement>
      </statements>
      <symbol>}</symbol>
    </subroutineBody>
  </subroutineDec>
  <symbol>}</symbol>
</class>`

			reader := strings.NewReader(inputTokensXml)
			testTokens := tokens.LoadTokensFromXML(reader)

			var buf bytes.Buffer
			xmlEnc := xml.NewEncoder(&buf)

			// Act
			ce, err := New(testTokens, xmlEnc)
			assert.NoError(t, err)
			err = ce.CompileClass()

			// Assert
			assert.NoError(t, err)

			assert.Equal(t, expectedXml, buf.String())
		})
		t.Run("should return XML when compiling Square.jack", func(t *testing.T) {

			// Arrange
			inputTokensXml := `<tokens>
  <keyword> class </keyword>
  <identifier> Square </identifier>
  <symbol> { </symbol>
  <keyword> field </keyword>
  <keyword> int </keyword>
  <identifier> x </identifier>
  <symbol> , </symbol>
  <identifier> y </identifier>
  <symbol> ; </symbol>
  <keyword> field </keyword>
  <keyword> int </keyword>
  <identifier> size </identifier>
  <symbol> ; </symbol>
  <keyword> constructor </keyword>
  <identifier> Square </identifier>
  <identifier> new </identifier>
  <symbol> ( </symbol>
  <keyword> int </keyword>
  <identifier> Ax </identifier>
  <symbol> , </symbol>
  <keyword> int </keyword>
  <identifier> Ay </identifier>
  <symbol> , </symbol>
  <keyword> int </keyword>
  <identifier> Asize </identifier>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> let </keyword>
  <identifier> x </identifier>
  <symbol> = </symbol>
  <identifier> Ax </identifier>
  <symbol> ; </symbol>
  <keyword> let </keyword>
  <identifier> y </identifier>
  <symbol> = </symbol>
  <identifier> Ay </identifier>
  <symbol> ; </symbol>
  <keyword> let </keyword>
  <identifier> size </identifier>
  <symbol> = </symbol>
  <identifier> Asize </identifier>
  <symbol> ; </symbol>
  <keyword> do </keyword>
  <identifier> draw </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> return </keyword>
  <keyword> this </keyword>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> method </keyword>
  <keyword> void </keyword>
  <identifier> dispose </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> do </keyword>
  <identifier> Memory </identifier>
  <symbol> . </symbol>
  <identifier> deAlloc </identifier>
  <symbol> ( </symbol>
  <keyword> this </keyword>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> return </keyword>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> method </keyword>
  <keyword> void </keyword>
  <identifier> draw </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> do </keyword>
  <identifier> Screen </identifier>
  <symbol> . </symbol>
  <identifier> setColor </identifier>
  <symbol> ( </symbol>
  <keyword> true </keyword>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> do </keyword>
  <identifier> Screen </identifier>
  <symbol> . </symbol>
  <identifier> drawRectangle </identifier>
  <symbol> ( </symbol>
  <identifier> x </identifier>
  <symbol> , </symbol>
  <identifier> y </identifier>
  <symbol> , </symbol>
  <identifier> x </identifier>
  <symbol> + </symbol>
  <identifier> size </identifier>
  <symbol> , </symbol>
  <identifier> y </identifier>
  <symbol> + </symbol>
  <identifier> size </identifier>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> return </keyword>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> method </keyword>
  <keyword> void </keyword>
  <identifier> erase </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> do </keyword>
  <identifier> Screen </identifier>
  <symbol> . </symbol>
  <identifier> setColor </identifier>
  <symbol> ( </symbol>
  <keyword> false </keyword>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> do </keyword>
  <identifier> Screen </identifier>
  <symbol> . </symbol>
  <identifier> drawRectangle </identifier>
  <symbol> ( </symbol>
  <identifier> x </identifier>
  <symbol> , </symbol>
  <identifier> y </identifier>
  <symbol> , </symbol>
  <identifier> x </identifier>
  <symbol> + </symbol>
  <identifier> size </identifier>
  <symbol> , </symbol>
  <identifier> y </identifier>
  <symbol> + </symbol>
  <identifier> size </identifier>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> return </keyword>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> method </keyword>
  <keyword> void </keyword>
  <identifier> incSize </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> if </keyword>
  <symbol> ( </symbol>
  <symbol> ( </symbol>
  <symbol> ( </symbol>
  <identifier> y </identifier>
  <symbol> + </symbol>
  <identifier> size </identifier>
  <symbol> ) </symbol>
  <symbol> &lt; </symbol>
  <integerConstant> 254 </integerConstant>
  <symbol> ) </symbol>
  <symbol> &amp; </symbol>
  <symbol> ( </symbol>
  <symbol> ( </symbol>
  <identifier> x </identifier>
  <symbol> + </symbol>
  <identifier> size </identifier>
  <symbol> ) </symbol>
  <symbol> &lt; </symbol>
  <integerConstant> 510 </integerConstant>
  <symbol> ) </symbol>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> do </keyword>
  <identifier> erase </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> let </keyword>
  <identifier> size </identifier>
  <symbol> = </symbol>
  <identifier> size </identifier>
  <symbol> + </symbol>
  <integerConstant> 2 </integerConstant>
  <symbol> ; </symbol>
  <keyword> do </keyword>
  <identifier> draw </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> return </keyword>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> method </keyword>
  <keyword> void </keyword>
  <identifier> decSize </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> if </keyword>
  <symbol> ( </symbol>
  <identifier> size </identifier>
  <symbol> &gt; </symbol>
  <integerConstant> 2 </integerConstant>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> do </keyword>
  <identifier> erase </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> let </keyword>
  <identifier> size </identifier>
  <symbol> = </symbol>
  <identifier> size </identifier>
  <symbol> - </symbol>
  <integerConstant> 2 </integerConstant>
  <symbol> ; </symbol>
  <keyword> do </keyword>
  <identifier> draw </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> return </keyword>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> method </keyword>
  <keyword> void </keyword>
  <identifier> moveUp </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> if </keyword>
  <symbol> ( </symbol>
  <identifier> y </identifier>
  <symbol> &gt; </symbol>
  <integerConstant> 1 </integerConstant>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> do </keyword>
  <identifier> Screen </identifier>
  <symbol> . </symbol>
  <identifier> setColor </identifier>
  <symbol> ( </symbol>
  <keyword> false </keyword>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> do </keyword>
  <identifier> Screen </identifier>
  <symbol> . </symbol>
  <identifier> drawRectangle </identifier>
  <symbol> ( </symbol>
  <identifier> x </identifier>
  <symbol> , </symbol>
  <symbol> ( </symbol>
  <identifier> y </identifier>
  <symbol> + </symbol>
  <identifier> size </identifier>
  <symbol> ) </symbol>
  <symbol> - </symbol>
  <integerConstant> 1 </integerConstant>
  <symbol> , </symbol>
  <identifier> x </identifier>
  <symbol> + </symbol>
  <identifier> size </identifier>
  <symbol> , </symbol>
  <identifier> y </identifier>
  <symbol> + </symbol>
  <identifier> size </identifier>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> let </keyword>
  <identifier> y </identifier>
  <symbol> = </symbol>
  <identifier> y </identifier>
  <symbol> - </symbol>
  <integerConstant> 2 </integerConstant>
  <symbol> ; </symbol>
  <keyword> do </keyword>
  <identifier> Screen </identifier>
  <symbol> . </symbol>
  <identifier> setColor </identifier>
  <symbol> ( </symbol>
  <keyword> true </keyword>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> do </keyword>
  <identifier> Screen </identifier>
  <symbol> . </symbol>
  <identifier> drawRectangle </identifier>
  <symbol> ( </symbol>
  <identifier> x </identifier>
  <symbol> , </symbol>
  <identifier> y </identifier>
  <symbol> , </symbol>
  <identifier> x </identifier>
  <symbol> + </symbol>
  <identifier> size </identifier>
  <symbol> , </symbol>
  <identifier> y </identifier>
  <symbol> + </symbol>
  <integerConstant> 1 </integerConstant>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> return </keyword>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> method </keyword>
  <keyword> void </keyword>
  <identifier> moveDown </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> if </keyword>
  <symbol> ( </symbol>
  <symbol> ( </symbol>
  <identifier> y </identifier>
  <symbol> + </symbol>
  <identifier> size </identifier>
  <symbol> ) </symbol>
  <symbol> &lt; </symbol>
  <integerConstant> 254 </integerConstant>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> do </keyword>
  <identifier> Screen </identifier>
  <symbol> . </symbol>
  <identifier> setColor </identifier>
  <symbol> ( </symbol>
  <keyword> false </keyword>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> do </keyword>
  <identifier> Screen </identifier>
  <symbol> . </symbol>
  <identifier> drawRectangle </identifier>
  <symbol> ( </symbol>
  <identifier> x </identifier>
  <symbol> , </symbol>
  <identifier> y </identifier>
  <symbol> , </symbol>
  <identifier> x </identifier>
  <symbol> + </symbol>
  <identifier> size </identifier>
  <symbol> , </symbol>
  <identifier> y </identifier>
  <symbol> + </symbol>
  <integerConstant> 1 </integerConstant>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> let </keyword>
  <identifier> y </identifier>
  <symbol> = </symbol>
  <identifier> y </identifier>
  <symbol> + </symbol>
  <integerConstant> 2 </integerConstant>
  <symbol> ; </symbol>
  <keyword> do </keyword>
  <identifier> Screen </identifier>
  <symbol> . </symbol>
  <identifier> setColor </identifier>
  <symbol> ( </symbol>
  <keyword> true </keyword>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> do </keyword>
  <identifier> Screen </identifier>
  <symbol> . </symbol>
  <identifier> drawRectangle </identifier>
  <symbol> ( </symbol>
  <identifier> x </identifier>
  <symbol> , </symbol>
  <symbol> ( </symbol>
  <identifier> y </identifier>
  <symbol> + </symbol>
  <identifier> size </identifier>
  <symbol> ) </symbol>
  <symbol> - </symbol>
  <integerConstant> 1 </integerConstant>
  <symbol> , </symbol>
  <identifier> x </identifier>
  <symbol> + </symbol>
  <identifier> size </identifier>
  <symbol> , </symbol>
  <identifier> y </identifier>
  <symbol> + </symbol>
  <identifier> size </identifier>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> return </keyword>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> method </keyword>
  <keyword> void </keyword>
  <identifier> moveLeft </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> if </keyword>
  <symbol> ( </symbol>
  <identifier> x </identifier>
  <symbol> &gt; </symbol>
  <integerConstant> 1 </integerConstant>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> do </keyword>
  <identifier> Screen </identifier>
  <symbol> . </symbol>
  <identifier> setColor </identifier>
  <symbol> ( </symbol>
  <keyword> false </keyword>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> do </keyword>
  <identifier> Screen </identifier>
  <symbol> . </symbol>
  <identifier> drawRectangle </identifier>
  <symbol> ( </symbol>
  <symbol> ( </symbol>
  <identifier> x </identifier>
  <symbol> + </symbol>
  <identifier> size </identifier>
  <symbol> ) </symbol>
  <symbol> - </symbol>
  <integerConstant> 1 </integerConstant>
  <symbol> , </symbol>
  <identifier> y </identifier>
  <symbol> , </symbol>
  <identifier> x </identifier>
  <symbol> + </symbol>
  <identifier> size </identifier>
  <symbol> , </symbol>
  <identifier> y </identifier>
  <symbol> + </symbol>
  <identifier> size </identifier>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> let </keyword>
  <identifier> x </identifier>
  <symbol> = </symbol>
  <identifier> x </identifier>
  <symbol> - </symbol>
  <integerConstant> 2 </integerConstant>
  <symbol> ; </symbol>
  <keyword> do </keyword>
  <identifier> Screen </identifier>
  <symbol> . </symbol>
  <identifier> setColor </identifier>
  <symbol> ( </symbol>
  <keyword> true </keyword>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> do </keyword>
  <identifier> Screen </identifier>
  <symbol> . </symbol>
  <identifier> drawRectangle </identifier>
  <symbol> ( </symbol>
  <identifier> x </identifier>
  <symbol> , </symbol>
  <identifier> y </identifier>
  <symbol> , </symbol>
  <identifier> x </identifier>
  <symbol> + </symbol>
  <integerConstant> 1 </integerConstant>
  <symbol> , </symbol>
  <identifier> y </identifier>
  <symbol> + </symbol>
  <identifier> size </identifier>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> return </keyword>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> method </keyword>
  <keyword> void </keyword>
  <identifier> moveRight </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> if </keyword>
  <symbol> ( </symbol>
  <symbol> ( </symbol>
  <identifier> x </identifier>
  <symbol> + </symbol>
  <identifier> size </identifier>
  <symbol> ) </symbol>
  <symbol> &lt; </symbol>
  <integerConstant> 510 </integerConstant>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> do </keyword>
  <identifier> Screen </identifier>
  <symbol> . </symbol>
  <identifier> setColor </identifier>
  <symbol> ( </symbol>
  <keyword> false </keyword>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> do </keyword>
  <identifier> Screen </identifier>
  <symbol> . </symbol>
  <identifier> drawRectangle </identifier>
  <symbol> ( </symbol>
  <identifier> x </identifier>
  <symbol> , </symbol>
  <identifier> y </identifier>
  <symbol> , </symbol>
  <identifier> x </identifier>
  <symbol> + </symbol>
  <integerConstant> 1 </integerConstant>
  <symbol> , </symbol>
  <identifier> y </identifier>
  <symbol> + </symbol>
  <identifier> size </identifier>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> let </keyword>
  <identifier> x </identifier>
  <symbol> = </symbol>
  <identifier> x </identifier>
  <symbol> + </symbol>
  <integerConstant> 2 </integerConstant>
  <symbol> ; </symbol>
  <keyword> do </keyword>
  <identifier> Screen </identifier>
  <symbol> . </symbol>
  <identifier> setColor </identifier>
  <symbol> ( </symbol>
  <keyword> true </keyword>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> do </keyword>
  <identifier> Screen </identifier>
  <symbol> . </symbol>
  <identifier> drawRectangle </identifier>
  <symbol> ( </symbol>
  <symbol> ( </symbol>
  <identifier> x </identifier>
  <symbol> + </symbol>
  <identifier> size </identifier>
  <symbol> ) </symbol>
  <symbol> - </symbol>
  <integerConstant> 1 </integerConstant>
  <symbol> , </symbol>
  <identifier> y </identifier>
  <symbol> , </symbol>
  <identifier> x </identifier>
  <symbol> + </symbol>
  <identifier> size </identifier>
  <symbol> , </symbol>
  <identifier> y </identifier>
  <symbol> + </symbol>
  <identifier> size </identifier>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> return </keyword>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <symbol> } </symbol>
</tokens>
`

			expectedXml := `<class>
  <keyword>class</keyword>
  <identifier>Square</identifier>
  <symbol>{</symbol>
  <classVarDec>
    <keyword>field</keyword>
    <keyword>int</keyword>
    <identifier>x</identifier>
    <symbol>,</symbol>
    <identifier>y</identifier>
    <symbol>;</symbol>
  </classVarDec>
  <classVarDec>
    <keyword>field</keyword>
    <keyword>int</keyword>
    <identifier>size</identifier>
    <symbol>;</symbol>
  </classVarDec>
  <subroutineDec>
    <keyword>constructor</keyword>
    <identifier>Square</identifier>
    <identifier>new</identifier>
    <symbol>(</symbol>
    <parameterList>
      <keyword>int</keyword>
      <identifier>Ax</identifier>
      <symbol>,</symbol>
      <keyword>int</keyword>
      <identifier>Ay</identifier>
      <symbol>,</symbol>
      <keyword>int</keyword>
      <identifier>Asize</identifier>
    </parameterList>
    <symbol>)</symbol>
    <subroutineBody>
      <symbol>{</symbol>
      <statements>
        <letStatement>
          <keyword>let</keyword>
          <identifier>x</identifier>
          <symbol>=</symbol>
          <expression>
            <term>
              <identifier>Ax</identifier>
            </term>
          </expression>
          <symbol>;</symbol>
        </letStatement>
        <letStatement>
          <keyword>let</keyword>
          <identifier>y</identifier>
          <symbol>=</symbol>
          <expression>
            <term>
              <identifier>Ay</identifier>
            </term>
          </expression>
          <symbol>;</symbol>
        </letStatement>
        <letStatement>
          <keyword>let</keyword>
          <identifier>size</identifier>
          <symbol>=</symbol>
          <expression>
            <term>
              <identifier>Asize</identifier>
            </term>
          </expression>
          <symbol>;</symbol>
        </letStatement>
        <doStatement>
          <keyword>do</keyword>
          <identifier>draw</identifier>
          <symbol>(</symbol>
          <expressionList></expressionList>
          <symbol>)</symbol>
          <symbol>;</symbol>
        </doStatement>
        <returnStatement>
          <keyword>return</keyword>
          <expression>
            <term>
              <keyword>this</keyword>
            </term>
          </expression>
          <symbol>;</symbol>
        </returnStatement>
      </statements>
      <symbol>}</symbol>
    </subroutineBody>
  </subroutineDec>
  <subroutineDec>
    <keyword>method</keyword>
    <keyword>void</keyword>
    <identifier>dispose</identifier>
    <symbol>(</symbol>
    <parameterList></parameterList>
    <symbol>)</symbol>
    <subroutineBody>
      <symbol>{</symbol>
      <statements>
        <doStatement>
          <keyword>do</keyword>
          <identifier>Memory</identifier>
          <symbol>.</symbol>
          <identifier>deAlloc</identifier>
          <symbol>(</symbol>
          <expressionList>
            <expression>
              <term>
                <keyword>this</keyword>
              </term>
            </expression>
          </expressionList>
          <symbol>)</symbol>
          <symbol>;</symbol>
        </doStatement>
        <returnStatement>
          <keyword>return</keyword>
          <symbol>;</symbol>
        </returnStatement>
      </statements>
      <symbol>}</symbol>
    </subroutineBody>
  </subroutineDec>
  <subroutineDec>
    <keyword>method</keyword>
    <keyword>void</keyword>
    <identifier>draw</identifier>
    <symbol>(</symbol>
    <parameterList></parameterList>
    <symbol>)</symbol>
    <subroutineBody>
      <symbol>{</symbol>
      <statements>
        <doStatement>
          <keyword>do</keyword>
          <identifier>Screen</identifier>
          <symbol>.</symbol>
          <identifier>setColor</identifier>
          <symbol>(</symbol>
          <expressionList>
            <expression>
              <term>
                <keyword>true</keyword>
              </term>
            </expression>
          </expressionList>
          <symbol>)</symbol>
          <symbol>;</symbol>
        </doStatement>
        <doStatement>
          <keyword>do</keyword>
          <identifier>Screen</identifier>
          <symbol>.</symbol>
          <identifier>drawRectangle</identifier>
          <symbol>(</symbol>
          <expressionList>
            <expression>
              <term>
                <identifier>x</identifier>
              </term>
            </expression>
            <symbol>,</symbol>
            <expression>
              <term>
                <identifier>y</identifier>
              </term>
            </expression>
            <symbol>,</symbol>
            <expression>
              <term>
                <identifier>x</identifier>
              </term>
              <symbol>+</symbol>
              <term>
                <identifier>size</identifier>
              </term>
            </expression>
            <symbol>,</symbol>
            <expression>
              <term>
                <identifier>y</identifier>
              </term>
              <symbol>+</symbol>
              <term>
                <identifier>size</identifier>
              </term>
            </expression>
          </expressionList>
          <symbol>)</symbol>
          <symbol>;</symbol>
        </doStatement>
        <returnStatement>
          <keyword>return</keyword>
          <symbol>;</symbol>
        </returnStatement>
      </statements>
      <symbol>}</symbol>
    </subroutineBody>
  </subroutineDec>
  <subroutineDec>
    <keyword>method</keyword>
    <keyword>void</keyword>
    <identifier>erase</identifier>
    <symbol>(</symbol>
    <parameterList></parameterList>
    <symbol>)</symbol>
    <subroutineBody>
      <symbol>{</symbol>
      <statements>
        <doStatement>
          <keyword>do</keyword>
          <identifier>Screen</identifier>
          <symbol>.</symbol>
          <identifier>setColor</identifier>
          <symbol>(</symbol>
          <expressionList>
            <expression>
              <term>
                <keyword>false</keyword>
              </term>
            </expression>
          </expressionList>
          <symbol>)</symbol>
          <symbol>;</symbol>
        </doStatement>
        <doStatement>
          <keyword>do</keyword>
          <identifier>Screen</identifier>
          <symbol>.</symbol>
          <identifier>drawRectangle</identifier>
          <symbol>(</symbol>
          <expressionList>
            <expression>
              <term>
                <identifier>x</identifier>
              </term>
            </expression>
            <symbol>,</symbol>
            <expression>
              <term>
                <identifier>y</identifier>
              </term>
            </expression>
            <symbol>,</symbol>
            <expression>
              <term>
                <identifier>x</identifier>
              </term>
              <symbol>+</symbol>
              <term>
                <identifier>size</identifier>
              </term>
            </expression>
            <symbol>,</symbol>
            <expression>
              <term>
                <identifier>y</identifier>
              </term>
              <symbol>+</symbol>
              <term>
                <identifier>size</identifier>
              </term>
            </expression>
          </expressionList>
          <symbol>)</symbol>
          <symbol>;</symbol>
        </doStatement>
        <returnStatement>
          <keyword>return</keyword>
          <symbol>;</symbol>
        </returnStatement>
      </statements>
      <symbol>}</symbol>
    </subroutineBody>
  </subroutineDec>
  <subroutineDec>
    <keyword>method</keyword>
    <keyword>void</keyword>
    <identifier>incSize</identifier>
    <symbol>(</symbol>
    <parameterList></parameterList>
    <symbol>)</symbol>
    <subroutineBody>
      <symbol>{</symbol>
      <statements>
        <ifStatement>
          <keyword>if</keyword>
          <symbol>(</symbol>
          <expression>
            <term>
              <symbol>(</symbol>
              <expression>
                <term>
                  <symbol>(</symbol>
                  <expression>
                    <term>
                      <identifier>y</identifier>
                    </term>
                    <symbol>+</symbol>
                    <term>
                      <identifier>size</identifier>
                    </term>
                  </expression>
                  <symbol>)</symbol>
                </term>
                <symbol>&lt;</symbol>
                <term>
                  <integerConstant>254</integerConstant>
                </term>
              </expression>
              <symbol>)</symbol>
            </term>
            <symbol>&amp;</symbol>
            <term>
              <symbol>(</symbol>
              <expression>
                <term>
                  <symbol>(</symbol>
                  <expression>
                    <term>
                      <identifier>x</identifier>
                    </term>
                    <symbol>+</symbol>
                    <term>
                      <identifier>size</identifier>
                    </term>
                  </expression>
                  <symbol>)</symbol>
                </term>
                <symbol>&lt;</symbol>
                <term>
                  <integerConstant>510</integerConstant>
                </term>
              </expression>
              <symbol>)</symbol>
            </term>
          </expression>
          <symbol>)</symbol>
          <symbol>{</symbol>
          <statements>
            <doStatement>
              <keyword>do</keyword>
              <identifier>erase</identifier>
              <symbol>(</symbol>
              <expressionList></expressionList>
              <symbol>)</symbol>
              <symbol>;</symbol>
            </doStatement>
            <letStatement>
              <keyword>let</keyword>
              <identifier>size</identifier>
              <symbol>=</symbol>
              <expression>
                <term>
                  <identifier>size</identifier>
                </term>
                <symbol>+</symbol>
                <term>
                  <integerConstant>2</integerConstant>
                </term>
              </expression>
              <symbol>;</symbol>
            </letStatement>
            <doStatement>
              <keyword>do</keyword>
              <identifier>draw</identifier>
              <symbol>(</symbol>
              <expressionList></expressionList>
              <symbol>)</symbol>
              <symbol>;</symbol>
            </doStatement>
          </statements>
          <symbol>}</symbol>
        </ifStatement>
        <returnStatement>
          <keyword>return</keyword>
          <symbol>;</symbol>
        </returnStatement>
      </statements>
      <symbol>}</symbol>
    </subroutineBody>
  </subroutineDec>
  <subroutineDec>
    <keyword>method</keyword>
    <keyword>void</keyword>
    <identifier>decSize</identifier>
    <symbol>(</symbol>
    <parameterList></parameterList>
    <symbol>)</symbol>
    <subroutineBody>
      <symbol>{</symbol>
      <statements>
        <ifStatement>
          <keyword>if</keyword>
          <symbol>(</symbol>
          <expression>
            <term>
              <identifier>size</identifier>
            </term>
            <symbol>&gt;</symbol>
            <term>
              <integerConstant>2</integerConstant>
            </term>
          </expression>
          <symbol>)</symbol>
          <symbol>{</symbol>
          <statements>
            <doStatement>
              <keyword>do</keyword>
              <identifier>erase</identifier>
              <symbol>(</symbol>
              <expressionList></expressionList>
              <symbol>)</symbol>
              <symbol>;</symbol>
            </doStatement>
            <letStatement>
              <keyword>let</keyword>
              <identifier>size</identifier>
              <symbol>=</symbol>
              <expression>
                <term>
                  <identifier>size</identifier>
                </term>
                <symbol>-</symbol>
                <term>
                  <integerConstant>2</integerConstant>
                </term>
              </expression>
              <symbol>;</symbol>
            </letStatement>
            <doStatement>
              <keyword>do</keyword>
              <identifier>draw</identifier>
              <symbol>(</symbol>
              <expressionList></expressionList>
              <symbol>)</symbol>
              <symbol>;</symbol>
            </doStatement>
          </statements>
          <symbol>}</symbol>
        </ifStatement>
        <returnStatement>
          <keyword>return</keyword>
          <symbol>;</symbol>
        </returnStatement>
      </statements>
      <symbol>}</symbol>
    </subroutineBody>
  </subroutineDec>
  <subroutineDec>
    <keyword>method</keyword>
    <keyword>void</keyword>
    <identifier>moveUp</identifier>
    <symbol>(</symbol>
    <parameterList></parameterList>
    <symbol>)</symbol>
    <subroutineBody>
      <symbol>{</symbol>
      <statements>
        <ifStatement>
          <keyword>if</keyword>
          <symbol>(</symbol>
          <expression>
            <term>
              <identifier>y</identifier>
            </term>
            <symbol>&gt;</symbol>
            <term>
              <integerConstant>1</integerConstant>
            </term>
          </expression>
          <symbol>)</symbol>
          <symbol>{</symbol>
          <statements>
            <doStatement>
              <keyword>do</keyword>
              <identifier>Screen</identifier>
              <symbol>.</symbol>
              <identifier>setColor</identifier>
              <symbol>(</symbol>
              <expressionList>
                <expression>
                  <term>
                    <keyword>false</keyword>
                  </term>
                </expression>
              </expressionList>
              <symbol>)</symbol>
              <symbol>;</symbol>
            </doStatement>
            <doStatement>
              <keyword>do</keyword>
              <identifier>Screen</identifier>
              <symbol>.</symbol>
              <identifier>drawRectangle</identifier>
              <symbol>(</symbol>
              <expressionList>
                <expression>
                  <term>
                    <identifier>x</identifier>
                  </term>
                </expression>
                <symbol>,</symbol>
                <expression>
                  <term>
                    <symbol>(</symbol>
                    <expression>
                      <term>
                        <identifier>y</identifier>
                      </term>
                      <symbol>+</symbol>
                      <term>
                        <identifier>size</identifier>
                      </term>
                    </expression>
                    <symbol>)</symbol>
                  </term>
                  <symbol>-</symbol>
                  <term>
                    <integerConstant>1</integerConstant>
                  </term>
                </expression>
                <symbol>,</symbol>
                <expression>
                  <term>
                    <identifier>x</identifier>
                  </term>
                  <symbol>+</symbol>
                  <term>
                    <identifier>size</identifier>
                  </term>
                </expression>
                <symbol>,</symbol>
                <expression>
                  <term>
                    <identifier>y</identifier>
                  </term>
                  <symbol>+</symbol>
                  <term>
                    <identifier>size</identifier>
                  </term>
                </expression>
              </expressionList>
              <symbol>)</symbol>
              <symbol>;</symbol>
            </doStatement>
            <letStatement>
              <keyword>let</keyword>
              <identifier>y</identifier>
              <symbol>=</symbol>
              <expression>
                <term>
                  <identifier>y</identifier>
                </term>
                <symbol>-</symbol>
                <term>
                  <integerConstant>2</integerConstant>
                </term>
              </expression>
              <symbol>;</symbol>
            </letStatement>
            <doStatement>
              <keyword>do</keyword>
              <identifier>Screen</identifier>
              <symbol>.</symbol>
              <identifier>setColor</identifier>
              <symbol>(</symbol>
              <expressionList>
                <expression>
                  <term>
                    <keyword>true</keyword>
                  </term>
                </expression>
              </expressionList>
              <symbol>)</symbol>
              <symbol>;</symbol>
            </doStatement>
            <doStatement>
              <keyword>do</keyword>
              <identifier>Screen</identifier>
              <symbol>.</symbol>
              <identifier>drawRectangle</identifier>
              <symbol>(</symbol>
              <expressionList>
                <expression>
                  <term>
                    <identifier>x</identifier>
                  </term>
                </expression>
                <symbol>,</symbol>
                <expression>
                  <term>
                    <identifier>y</identifier>
                  </term>
                </expression>
                <symbol>,</symbol>
                <expression>
                  <term>
                    <identifier>x</identifier>
                  </term>
                  <symbol>+</symbol>
                  <term>
                    <identifier>size</identifier>
                  </term>
                </expression>
                <symbol>,</symbol>
                <expression>
                  <term>
                    <identifier>y</identifier>
                  </term>
                  <symbol>+</symbol>
                  <term>
                    <integerConstant>1</integerConstant>
                  </term>
                </expression>
              </expressionList>
              <symbol>)</symbol>
              <symbol>;</symbol>
            </doStatement>
          </statements>
          <symbol>}</symbol>
        </ifStatement>
        <returnStatement>
          <keyword>return</keyword>
          <symbol>;</symbol>
        </returnStatement>
      </statements>
      <symbol>}</symbol>
    </subroutineBody>
  </subroutineDec>
  <subroutineDec>
    <keyword>method</keyword>
    <keyword>void</keyword>
    <identifier>moveDown</identifier>
    <symbol>(</symbol>
    <parameterList></parameterList>
    <symbol>)</symbol>
    <subroutineBody>
      <symbol>{</symbol>
      <statements>
        <ifStatement>
          <keyword>if</keyword>
          <symbol>(</symbol>
          <expression>
            <term>
              <symbol>(</symbol>
              <expression>
                <term>
                  <identifier>y</identifier>
                </term>
                <symbol>+</symbol>
                <term>
                  <identifier>size</identifier>
                </term>
              </expression>
              <symbol>)</symbol>
            </term>
            <symbol>&lt;</symbol>
            <term>
              <integerConstant>254</integerConstant>
            </term>
          </expression>
          <symbol>)</symbol>
          <symbol>{</symbol>
          <statements>
            <doStatement>
              <keyword>do</keyword>
              <identifier>Screen</identifier>
              <symbol>.</symbol>
              <identifier>setColor</identifier>
              <symbol>(</symbol>
              <expressionList>
                <expression>
                  <term>
                    <keyword>false</keyword>
                  </term>
                </expression>
              </expressionList>
              <symbol>)</symbol>
              <symbol>;</symbol>
            </doStatement>
            <doStatement>
              <keyword>do</keyword>
              <identifier>Screen</identifier>
              <symbol>.</symbol>
              <identifier>drawRectangle</identifier>
              <symbol>(</symbol>
              <expressionList>
                <expression>
                  <term>
                    <identifier>x</identifier>
                  </term>
                </expression>
                <symbol>,</symbol>
                <expression>
                  <term>
                    <identifier>y</identifier>
                  </term>
                </expression>
                <symbol>,</symbol>
                <expression>
                  <term>
                    <identifier>x</identifier>
                  </term>
                  <symbol>+</symbol>
                  <term>
                    <identifier>size</identifier>
                  </term>
                </expression>
                <symbol>,</symbol>
                <expression>
                  <term>
                    <identifier>y</identifier>
                  </term>
                  <symbol>+</symbol>
                  <term>
                    <integerConstant>1</integerConstant>
                  </term>
                </expression>
              </expressionList>
              <symbol>)</symbol>
              <symbol>;</symbol>
            </doStatement>
            <letStatement>
              <keyword>let</keyword>
              <identifier>y</identifier>
              <symbol>=</symbol>
              <expression>
                <term>
                  <identifier>y</identifier>
                </term>
                <symbol>+</symbol>
                <term>
                  <integerConstant>2</integerConstant>
                </term>
              </expression>
              <symbol>;</symbol>
            </letStatement>
            <doStatement>
              <keyword>do</keyword>
              <identifier>Screen</identifier>
              <symbol>.</symbol>
              <identifier>setColor</identifier>
              <symbol>(</symbol>
              <expressionList>
                <expression>
                  <term>
                    <keyword>true</keyword>
                  </term>
                </expression>
              </expressionList>
              <symbol>)</symbol>
              <symbol>;</symbol>
            </doStatement>
            <doStatement>
              <keyword>do</keyword>
              <identifier>Screen</identifier>
              <symbol>.</symbol>
              <identifier>drawRectangle</identifier>
              <symbol>(</symbol>
              <expressionList>
                <expression>
                  <term>
                    <identifier>x</identifier>
                  </term>
                </expression>
                <symbol>,</symbol>
                <expression>
                  <term>
                    <symbol>(</symbol>
                    <expression>
                      <term>
                        <identifier>y</identifier>
                      </term>
                      <symbol>+</symbol>
                      <term>
                        <identifier>size</identifier>
                      </term>
                    </expression>
                    <symbol>)</symbol>
                  </term>
                  <symbol>-</symbol>
                  <term>
                    <integerConstant>1</integerConstant>
                  </term>
                </expression>
                <symbol>,</symbol>
                <expression>
                  <term>
                    <identifier>x</identifier>
                  </term>
                  <symbol>+</symbol>
                  <term>
                    <identifier>size</identifier>
                  </term>
                </expression>
                <symbol>,</symbol>
                <expression>
                  <term>
                    <identifier>y</identifier>
                  </term>
                  <symbol>+</symbol>
                  <term>
                    <identifier>size</identifier>
                  </term>
                </expression>
              </expressionList>
              <symbol>)</symbol>
              <symbol>;</symbol>
            </doStatement>
          </statements>
          <symbol>}</symbol>
        </ifStatement>
        <returnStatement>
          <keyword>return</keyword>
          <symbol>;</symbol>
        </returnStatement>
      </statements>
      <symbol>}</symbol>
    </subroutineBody>
  </subroutineDec>
  <subroutineDec>
    <keyword>method</keyword>
    <keyword>void</keyword>
    <identifier>moveLeft</identifier>
    <symbol>(</symbol>
    <parameterList></parameterList>
    <symbol>)</symbol>
    <subroutineBody>
      <symbol>{</symbol>
      <statements>
        <ifStatement>
          <keyword>if</keyword>
          <symbol>(</symbol>
          <expression>
            <term>
              <identifier>x</identifier>
            </term>
            <symbol>&gt;</symbol>
            <term>
              <integerConstant>1</integerConstant>
            </term>
          </expression>
          <symbol>)</symbol>
          <symbol>{</symbol>
          <statements>
            <doStatement>
              <keyword>do</keyword>
              <identifier>Screen</identifier>
              <symbol>.</symbol>
              <identifier>setColor</identifier>
              <symbol>(</symbol>
              <expressionList>
                <expression>
                  <term>
                    <keyword>false</keyword>
                  </term>
                </expression>
              </expressionList>
              <symbol>)</symbol>
              <symbol>;</symbol>
            </doStatement>
            <doStatement>
              <keyword>do</keyword>
              <identifier>Screen</identifier>
              <symbol>.</symbol>
              <identifier>drawRectangle</identifier>
              <symbol>(</symbol>
              <expressionList>
                <expression>
                  <term>
                    <symbol>(</symbol>
                    <expression>
                      <term>
                        <identifier>x</identifier>
                      </term>
                      <symbol>+</symbol>
                      <term>
                        <identifier>size</identifier>
                      </term>
                    </expression>
                    <symbol>)</symbol>
                  </term>
                  <symbol>-</symbol>
                  <term>
                    <integerConstant>1</integerConstant>
                  </term>
                </expression>
                <symbol>,</symbol>
                <expression>
                  <term>
                    <identifier>y</identifier>
                  </term>
                </expression>
                <symbol>,</symbol>
                <expression>
                  <term>
                    <identifier>x</identifier>
                  </term>
                  <symbol>+</symbol>
                  <term>
                    <identifier>size</identifier>
                  </term>
                </expression>
                <symbol>,</symbol>
                <expression>
                  <term>
                    <identifier>y</identifier>
                  </term>
                  <symbol>+</symbol>
                  <term>
                    <identifier>size</identifier>
                  </term>
                </expression>
              </expressionList>
              <symbol>)</symbol>
              <symbol>;</symbol>
            </doStatement>
            <letStatement>
              <keyword>let</keyword>
              <identifier>x</identifier>
              <symbol>=</symbol>
              <expression>
                <term>
                  <identifier>x</identifier>
                </term>
                <symbol>-</symbol>
                <term>
                  <integerConstant>2</integerConstant>
                </term>
              </expression>
              <symbol>;</symbol>
            </letStatement>
            <doStatement>
              <keyword>do</keyword>
              <identifier>Screen</identifier>
              <symbol>.</symbol>
              <identifier>setColor</identifier>
              <symbol>(</symbol>
              <expressionList>
                <expression>
                  <term>
                    <keyword>true</keyword>
                  </term>
                </expression>
              </expressionList>
              <symbol>)</symbol>
              <symbol>;</symbol>
            </doStatement>
            <doStatement>
              <keyword>do</keyword>
              <identifier>Screen</identifier>
              <symbol>.</symbol>
              <identifier>drawRectangle</identifier>
              <symbol>(</symbol>
              <expressionList>
                <expression>
                  <term>
                    <identifier>x</identifier>
                  </term>
                </expression>
                <symbol>,</symbol>
                <expression>
                  <term>
                    <identifier>y</identifier>
                  </term>
                </expression>
                <symbol>,</symbol>
                <expression>
                  <term>
                    <identifier>x</identifier>
                  </term>
                  <symbol>+</symbol>
                  <term>
                    <integerConstant>1</integerConstant>
                  </term>
                </expression>
                <symbol>,</symbol>
                <expression>
                  <term>
                    <identifier>y</identifier>
                  </term>
                  <symbol>+</symbol>
                  <term>
                    <identifier>size</identifier>
                  </term>
                </expression>
              </expressionList>
              <symbol>)</symbol>
              <symbol>;</symbol>
            </doStatement>
          </statements>
          <symbol>}</symbol>
        </ifStatement>
        <returnStatement>
          <keyword>return</keyword>
          <symbol>;</symbol>
        </returnStatement>
      </statements>
      <symbol>}</symbol>
    </subroutineBody>
  </subroutineDec>
  <subroutineDec>
    <keyword>method</keyword>
    <keyword>void</keyword>
    <identifier>moveRight</identifier>
    <symbol>(</symbol>
    <parameterList></parameterList>
    <symbol>)</symbol>
    <subroutineBody>
      <symbol>{</symbol>
      <statements>
        <ifStatement>
          <keyword>if</keyword>
          <symbol>(</symbol>
          <expression>
            <term>
              <symbol>(</symbol>
              <expression>
                <term>
                  <identifier>x</identifier>
                </term>
                <symbol>+</symbol>
                <term>
                  <identifier>size</identifier>
                </term>
              </expression>
              <symbol>)</symbol>
            </term>
            <symbol>&lt;</symbol>
            <term>
              <integerConstant>510</integerConstant>
            </term>
          </expression>
          <symbol>)</symbol>
          <symbol>{</symbol>
          <statements>
            <doStatement>
              <keyword>do</keyword>
              <identifier>Screen</identifier>
              <symbol>.</symbol>
              <identifier>setColor</identifier>
              <symbol>(</symbol>
              <expressionList>
                <expression>
                  <term>
                    <keyword>false</keyword>
                  </term>
                </expression>
              </expressionList>
              <symbol>)</symbol>
              <symbol>;</symbol>
            </doStatement>
            <doStatement>
              <keyword>do</keyword>
              <identifier>Screen</identifier>
              <symbol>.</symbol>
              <identifier>drawRectangle</identifier>
              <symbol>(</symbol>
              <expressionList>
                <expression>
                  <term>
                    <identifier>x</identifier>
                  </term>
                </expression>
                <symbol>,</symbol>
                <expression>
                  <term>
                    <identifier>y</identifier>
                  </term>
                </expression>
                <symbol>,</symbol>
                <expression>
                  <term>
                    <identifier>x</identifier>
                  </term>
                  <symbol>+</symbol>
                  <term>
                    <integerConstant>1</integerConstant>
                  </term>
                </expression>
                <symbol>,</symbol>
                <expression>
                  <term>
                    <identifier>y</identifier>
                  </term>
                  <symbol>+</symbol>
                  <term>
                    <identifier>size</identifier>
                  </term>
                </expression>
              </expressionList>
              <symbol>)</symbol>
              <symbol>;</symbol>
            </doStatement>
            <letStatement>
              <keyword>let</keyword>
              <identifier>x</identifier>
              <symbol>=</symbol>
              <expression>
                <term>
                  <identifier>x</identifier>
                </term>
                <symbol>+</symbol>
                <term>
                  <integerConstant>2</integerConstant>
                </term>
              </expression>
              <symbol>;</symbol>
            </letStatement>
            <doStatement>
              <keyword>do</keyword>
              <identifier>Screen</identifier>
              <symbol>.</symbol>
              <identifier>setColor</identifier>
              <symbol>(</symbol>
              <expressionList>
                <expression>
                  <term>
                    <keyword>true</keyword>
                  </term>
                </expression>
              </expressionList>
              <symbol>)</symbol>
              <symbol>;</symbol>
            </doStatement>
            <doStatement>
              <keyword>do</keyword>
              <identifier>Screen</identifier>
              <symbol>.</symbol>
              <identifier>drawRectangle</identifier>
              <symbol>(</symbol>
              <expressionList>
                <expression>
                  <term>
                    <symbol>(</symbol>
                    <expression>
                      <term>
                        <identifier>x</identifier>
                      </term>
                      <symbol>+</symbol>
                      <term>
                        <identifier>size</identifier>
                      </term>
                    </expression>
                    <symbol>)</symbol>
                  </term>
                  <symbol>-</symbol>
                  <term>
                    <integerConstant>1</integerConstant>
                  </term>
                </expression>
                <symbol>,</symbol>
                <expression>
                  <term>
                    <identifier>y</identifier>
                  </term>
                </expression>
                <symbol>,</symbol>
                <expression>
                  <term>
                    <identifier>x</identifier>
                  </term>
                  <symbol>+</symbol>
                  <term>
                    <identifier>size</identifier>
                  </term>
                </expression>
                <symbol>,</symbol>
                <expression>
                  <term>
                    <identifier>y</identifier>
                  </term>
                  <symbol>+</symbol>
                  <term>
                    <identifier>size</identifier>
                  </term>
                </expression>
              </expressionList>
              <symbol>)</symbol>
              <symbol>;</symbol>
            </doStatement>
          </statements>
          <symbol>}</symbol>
        </ifStatement>
        <returnStatement>
          <keyword>return</keyword>
          <symbol>;</symbol>
        </returnStatement>
      </statements>
      <symbol>}</symbol>
    </subroutineBody>
  </subroutineDec>
  <symbol>}</symbol>
</class>`

			reader := strings.NewReader(inputTokensXml)
			testTokens := tokens.LoadTokensFromXML(reader)

			var buf bytes.Buffer
			xmlEnc := xml.NewEncoder(&buf)

			// Act
			ce, err := New(testTokens, xmlEnc)
			assert.NoError(t, err)
			err = ce.CompileClass()

			// Assert
			assert.NoError(t, err)

			assert.Equal(t, expectedXml, buf.String())
		})

		t.Run("should return XML when compiling SquareGame.jack", func(t *testing.T) {

			// Arrange
			inputTokensXml := `<tokens>
  <keyword> class </keyword>
  <identifier> SquareGame </identifier>
  <symbol> { </symbol>
  <keyword> field </keyword>
  <identifier> Square </identifier>
  <identifier> square </identifier>
  <symbol> ; </symbol>
  <keyword> field </keyword>
  <keyword> int </keyword>
  <identifier> direction </identifier>
  <symbol> ; </symbol>
  <keyword> constructor </keyword>
  <identifier> SquareGame </identifier>
  <identifier> new </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> let </keyword>
  <identifier> square </identifier>
  <symbol> = </symbol>
  <identifier> Square </identifier>
  <symbol> . </symbol>
  <identifier> new </identifier>
  <symbol> ( </symbol>
  <integerConstant> 0 </integerConstant>
  <symbol> , </symbol>
  <integerConstant> 0 </integerConstant>
  <symbol> , </symbol>
  <integerConstant> 30 </integerConstant>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> let </keyword>
  <identifier> direction </identifier>
  <symbol> = </symbol>
  <integerConstant> 0 </integerConstant>
  <symbol> ; </symbol>
  <keyword> return </keyword>
  <keyword> this </keyword>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> method </keyword>
  <keyword> void </keyword>
  <identifier> dispose </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> do </keyword>
  <identifier> square </identifier>
  <symbol> . </symbol>
  <identifier> dispose </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> do </keyword>
  <identifier> Memory </identifier>
  <symbol> . </symbol>
  <identifier> deAlloc </identifier>
  <symbol> ( </symbol>
  <keyword> this </keyword>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> return </keyword>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> method </keyword>
  <keyword> void </keyword>
  <identifier> moveSquare </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> if </keyword>
  <symbol> ( </symbol>
  <identifier> direction </identifier>
  <symbol> = </symbol>
  <integerConstant> 1 </integerConstant>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> do </keyword>
  <identifier> square </identifier>
  <symbol> . </symbol>
  <identifier> moveUp </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> if </keyword>
  <symbol> ( </symbol>
  <identifier> direction </identifier>
  <symbol> = </symbol>
  <integerConstant> 2 </integerConstant>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> do </keyword>
  <identifier> square </identifier>
  <symbol> . </symbol>
  <identifier> moveDown </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> if </keyword>
  <symbol> ( </symbol>
  <identifier> direction </identifier>
  <symbol> = </symbol>
  <integerConstant> 3 </integerConstant>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> do </keyword>
  <identifier> square </identifier>
  <symbol> . </symbol>
  <identifier> moveLeft </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> if </keyword>
  <symbol> ( </symbol>
  <identifier> direction </identifier>
  <symbol> = </symbol>
  <integerConstant> 4 </integerConstant>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> do </keyword>
  <identifier> square </identifier>
  <symbol> . </symbol>
  <identifier> moveRight </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> do </keyword>
  <identifier> Sys </identifier>
  <symbol> . </symbol>
  <identifier> wait </identifier>
  <symbol> ( </symbol>
  <integerConstant> 5 </integerConstant>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> return </keyword>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> method </keyword>
  <keyword> void </keyword>
  <identifier> run </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> var </keyword>
  <keyword> char </keyword>
  <identifier> key </identifier>
  <symbol> ; </symbol>
  <keyword> var </keyword>
  <keyword> boolean </keyword>
  <identifier> exit </identifier>
  <symbol> ; </symbol>
  <keyword> let </keyword>
  <identifier> exit </identifier>
  <symbol> = </symbol>
  <keyword> false </keyword>
  <symbol> ; </symbol>
  <keyword> while </keyword>
  <symbol> ( </symbol>
  <symbol> ~ </symbol>
  <identifier> exit </identifier>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> while </keyword>
  <symbol> ( </symbol>
  <identifier> key </identifier>
  <symbol> = </symbol>
  <integerConstant> 0 </integerConstant>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> let </keyword>
  <identifier> key </identifier>
  <symbol> = </symbol>
  <identifier> Keyboard </identifier>
  <symbol> . </symbol>
  <identifier> keyPressed </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> do </keyword>
  <identifier> moveSquare </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> if </keyword>
  <symbol> ( </symbol>
  <identifier> key </identifier>
  <symbol> = </symbol>
  <integerConstant> 81 </integerConstant>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> let </keyword>
  <identifier> exit </identifier>
  <symbol> = </symbol>
  <keyword> true </keyword>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> if </keyword>
  <symbol> ( </symbol>
  <identifier> key </identifier>
  <symbol> = </symbol>
  <integerConstant> 90 </integerConstant>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> do </keyword>
  <identifier> square </identifier>
  <symbol> . </symbol>
  <identifier> decSize </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> if </keyword>
  <symbol> ( </symbol>
  <identifier> key </identifier>
  <symbol> = </symbol>
  <integerConstant> 88 </integerConstant>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> do </keyword>
  <identifier> square </identifier>
  <symbol> . </symbol>
  <identifier> incSize </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> if </keyword>
  <symbol> ( </symbol>
  <identifier> key </identifier>
  <symbol> = </symbol>
  <integerConstant> 131 </integerConstant>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> let </keyword>
  <identifier> direction </identifier>
  <symbol> = </symbol>
  <integerConstant> 1 </integerConstant>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> if </keyword>
  <symbol> ( </symbol>
  <identifier> key </identifier>
  <symbol> = </symbol>
  <integerConstant> 133 </integerConstant>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> let </keyword>
  <identifier> direction </identifier>
  <symbol> = </symbol>
  <integerConstant> 2 </integerConstant>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> if </keyword>
  <symbol> ( </symbol>
  <identifier> key </identifier>
  <symbol> = </symbol>
  <integerConstant> 130 </integerConstant>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> let </keyword>
  <identifier> direction </identifier>
  <symbol> = </symbol>
  <integerConstant> 3 </integerConstant>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> if </keyword>
  <symbol> ( </symbol>
  <identifier> key </identifier>
  <symbol> = </symbol>
  <integerConstant> 132 </integerConstant>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> let </keyword>
  <identifier> direction </identifier>
  <symbol> = </symbol>
  <integerConstant> 4 </integerConstant>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <keyword> while </keyword>
  <symbol> ( </symbol>
  <symbol> ~ </symbol>
  <symbol> ( </symbol>
  <identifier> key </identifier>
  <symbol> = </symbol>
  <integerConstant> 0 </integerConstant>
  <symbol> ) </symbol>
  <symbol> ) </symbol>
  <symbol> { </symbol>
  <keyword> let </keyword>
  <identifier> key </identifier>
  <symbol> = </symbol>
  <identifier> Keyboard </identifier>
  <symbol> . </symbol>
  <identifier> keyPressed </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <keyword> do </keyword>
  <identifier> moveSquare </identifier>
  <symbol> ( </symbol>
  <symbol> ) </symbol>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <symbol> } </symbol>
  <keyword> return </keyword>
  <symbol> ; </symbol>
  <symbol> } </symbol>
  <symbol> } </symbol>
</tokens>
`

			expectedXml := `<class>
  <keyword>class</keyword>
  <identifier>SquareGame</identifier>
  <symbol>{</symbol>
  <classVarDec>
    <keyword>field</keyword>
    <identifier>Square</identifier>
    <identifier>square</identifier>
    <symbol>;</symbol>
  </classVarDec>
  <classVarDec>
    <keyword>field</keyword>
    <keyword>int</keyword>
    <identifier>direction</identifier>
    <symbol>;</symbol>
  </classVarDec>
  <subroutineDec>
    <keyword>constructor</keyword>
    <identifier>SquareGame</identifier>
    <identifier>new</identifier>
    <symbol>(</symbol>
    <parameterList></parameterList>
    <symbol>)</symbol>
    <subroutineBody>
      <symbol>{</symbol>
      <statements>
        <letStatement>
          <keyword>let</keyword>
          <identifier>square</identifier>
          <symbol>=</symbol>
          <expression>
            <term>
              <identifier>Square</identifier>
              <symbol>.</symbol>
              <identifier>new</identifier>
              <symbol>(</symbol>
              <expressionList>
                <expression>
                  <term>
                    <integerConstant>0</integerConstant>
                  </term>
                </expression>
                <symbol>,</symbol>
                <expression>
                  <term>
                    <integerConstant>0</integerConstant>
                  </term>
                </expression>
                <symbol>,</symbol>
                <expression>
                  <term>
                    <integerConstant>30</integerConstant>
                  </term>
                </expression>
              </expressionList>
              <symbol>)</symbol>
            </term>
          </expression>
          <symbol>;</symbol>
        </letStatement>
        <letStatement>
          <keyword>let</keyword>
          <identifier>direction</identifier>
          <symbol>=</symbol>
          <expression>
            <term>
              <integerConstant>0</integerConstant>
            </term>
          </expression>
          <symbol>;</symbol>
        </letStatement>
        <returnStatement>
          <keyword>return</keyword>
          <expression>
            <term>
              <keyword>this</keyword>
            </term>
          </expression>
          <symbol>;</symbol>
        </returnStatement>
      </statements>
      <symbol>}</symbol>
    </subroutineBody>
  </subroutineDec>
  <subroutineDec>
    <keyword>method</keyword>
    <keyword>void</keyword>
    <identifier>dispose</identifier>
    <symbol>(</symbol>
    <parameterList></parameterList>
    <symbol>)</symbol>
    <subroutineBody>
      <symbol>{</symbol>
      <statements>
        <doStatement>
          <keyword>do</keyword>
          <identifier>square</identifier>
          <symbol>.</symbol>
          <identifier>dispose</identifier>
          <symbol>(</symbol>
          <expressionList></expressionList>
          <symbol>)</symbol>
          <symbol>;</symbol>
        </doStatement>
        <doStatement>
          <keyword>do</keyword>
          <identifier>Memory</identifier>
          <symbol>.</symbol>
          <identifier>deAlloc</identifier>
          <symbol>(</symbol>
          <expressionList>
            <expression>
              <term>
                <keyword>this</keyword>
              </term>
            </expression>
          </expressionList>
          <symbol>)</symbol>
          <symbol>;</symbol>
        </doStatement>
        <returnStatement>
          <keyword>return</keyword>
          <symbol>;</symbol>
        </returnStatement>
      </statements>
      <symbol>}</symbol>
    </subroutineBody>
  </subroutineDec>
  <subroutineDec>
    <keyword>method</keyword>
    <keyword>void</keyword>
    <identifier>moveSquare</identifier>
    <symbol>(</symbol>
    <parameterList></parameterList>
    <symbol>)</symbol>
    <subroutineBody>
      <symbol>{</symbol>
      <statements>
        <ifStatement>
          <keyword>if</keyword>
          <symbol>(</symbol>
          <expression>
            <term>
              <identifier>direction</identifier>
            </term>
            <symbol>=</symbol>
            <term>
              <integerConstant>1</integerConstant>
            </term>
          </expression>
          <symbol>)</symbol>
          <symbol>{</symbol>
          <statements>
            <doStatement>
              <keyword>do</keyword>
              <identifier>square</identifier>
              <symbol>.</symbol>
              <identifier>moveUp</identifier>
              <symbol>(</symbol>
              <expressionList></expressionList>
              <symbol>)</symbol>
              <symbol>;</symbol>
            </doStatement>
          </statements>
          <symbol>}</symbol>
        </ifStatement>
        <ifStatement>
          <keyword>if</keyword>
          <symbol>(</symbol>
          <expression>
            <term>
              <identifier>direction</identifier>
            </term>
            <symbol>=</symbol>
            <term>
              <integerConstant>2</integerConstant>
            </term>
          </expression>
          <symbol>)</symbol>
          <symbol>{</symbol>
          <statements>
            <doStatement>
              <keyword>do</keyword>
              <identifier>square</identifier>
              <symbol>.</symbol>
              <identifier>moveDown</identifier>
              <symbol>(</symbol>
              <expressionList></expressionList>
              <symbol>)</symbol>
              <symbol>;</symbol>
            </doStatement>
          </statements>
          <symbol>}</symbol>
        </ifStatement>
        <ifStatement>
          <keyword>if</keyword>
          <symbol>(</symbol>
          <expression>
            <term>
              <identifier>direction</identifier>
            </term>
            <symbol>=</symbol>
            <term>
              <integerConstant>3</integerConstant>
            </term>
          </expression>
          <symbol>)</symbol>
          <symbol>{</symbol>
          <statements>
            <doStatement>
              <keyword>do</keyword>
              <identifier>square</identifier>
              <symbol>.</symbol>
              <identifier>moveLeft</identifier>
              <symbol>(</symbol>
              <expressionList></expressionList>
              <symbol>)</symbol>
              <symbol>;</symbol>
            </doStatement>
          </statements>
          <symbol>}</symbol>
        </ifStatement>
        <ifStatement>
          <keyword>if</keyword>
          <symbol>(</symbol>
          <expression>
            <term>
              <identifier>direction</identifier>
            </term>
            <symbol>=</symbol>
            <term>
              <integerConstant>4</integerConstant>
            </term>
          </expression>
          <symbol>)</symbol>
          <symbol>{</symbol>
          <statements>
            <doStatement>
              <keyword>do</keyword>
              <identifier>square</identifier>
              <symbol>.</symbol>
              <identifier>moveRight</identifier>
              <symbol>(</symbol>
              <expressionList></expressionList>
              <symbol>)</symbol>
              <symbol>;</symbol>
            </doStatement>
          </statements>
          <symbol>}</symbol>
        </ifStatement>
        <doStatement>
          <keyword>do</keyword>
          <identifier>Sys</identifier>
          <symbol>.</symbol>
          <identifier>wait</identifier>
          <symbol>(</symbol>
          <expressionList>
            <expression>
              <term>
                <integerConstant>5</integerConstant>
              </term>
            </expression>
          </expressionList>
          <symbol>)</symbol>
          <symbol>;</symbol>
        </doStatement>
        <returnStatement>
          <keyword>return</keyword>
          <symbol>;</symbol>
        </returnStatement>
      </statements>
      <symbol>}</symbol>
    </subroutineBody>
  </subroutineDec>
  <subroutineDec>
    <keyword>method</keyword>
    <keyword>void</keyword>
    <identifier>run</identifier>
    <symbol>(</symbol>
    <parameterList></parameterList>
    <symbol>)</symbol>
    <subroutineBody>
      <symbol>{</symbol>
      <varDec>
        <keyword>var</keyword>
        <keyword>char</keyword>
        <identifier>key</identifier>
        <symbol>;</symbol>
      </varDec>
      <varDec>
        <keyword>var</keyword>
        <keyword>boolean</keyword>
        <identifier>exit</identifier>
        <symbol>;</symbol>
      </varDec>
      <statements>
        <letStatement>
          <keyword>let</keyword>
          <identifier>exit</identifier>
          <symbol>=</symbol>
          <expression>
            <term>
              <keyword>false</keyword>
            </term>
          </expression>
          <symbol>;</symbol>
        </letStatement>
        <whileStatement>
          <keyword>while</keyword>
          <symbol>(</symbol>
          <expression>
            <term>
              <symbol>~</symbol>
              <term>
                <identifier>exit</identifier>
              </term>
            </term>
          </expression>
          <symbol>)</symbol>
          <symbol>{</symbol>
          <statements>
            <whileStatement>
              <keyword>while</keyword>
              <symbol>(</symbol>
              <expression>
                <term>
                  <identifier>key</identifier>
                </term>
                <symbol>=</symbol>
                <term>
                  <integerConstant>0</integerConstant>
                </term>
              </expression>
              <symbol>)</symbol>
              <symbol>{</symbol>
              <statements>
                <letStatement>
                  <keyword>let</keyword>
                  <identifier>key</identifier>
                  <symbol>=</symbol>
                  <expression>
                    <term>
                      <identifier>Keyboard</identifier>
                      <symbol>.</symbol>
                      <identifier>keyPressed</identifier>
                      <symbol>(</symbol>
                      <expressionList></expressionList>
                      <symbol>)</symbol>
                    </term>
                  </expression>
                  <symbol>;</symbol>
                </letStatement>
                <doStatement>
                  <keyword>do</keyword>
                  <identifier>moveSquare</identifier>
                  <symbol>(</symbol>
                  <expressionList></expressionList>
                  <symbol>)</symbol>
                  <symbol>;</symbol>
                </doStatement>
              </statements>
              <symbol>}</symbol>
            </whileStatement>
            <ifStatement>
              <keyword>if</keyword>
              <symbol>(</symbol>
              <expression>
                <term>
                  <identifier>key</identifier>
                </term>
                <symbol>=</symbol>
                <term>
                  <integerConstant>81</integerConstant>
                </term>
              </expression>
              <symbol>)</symbol>
              <symbol>{</symbol>
              <statements>
                <letStatement>
                  <keyword>let</keyword>
                  <identifier>exit</identifier>
                  <symbol>=</symbol>
                  <expression>
                    <term>
                      <keyword>true</keyword>
                    </term>
                  </expression>
                  <symbol>;</symbol>
                </letStatement>
              </statements>
              <symbol>}</symbol>
            </ifStatement>
            <ifStatement>
              <keyword>if</keyword>
              <symbol>(</symbol>
              <expression>
                <term>
                  <identifier>key</identifier>
                </term>
                <symbol>=</symbol>
                <term>
                  <integerConstant>90</integerConstant>
                </term>
              </expression>
              <symbol>)</symbol>
              <symbol>{</symbol>
              <statements>
                <doStatement>
                  <keyword>do</keyword>
                  <identifier>square</identifier>
                  <symbol>.</symbol>
                  <identifier>decSize</identifier>
                  <symbol>(</symbol>
                  <expressionList></expressionList>
                  <symbol>)</symbol>
                  <symbol>;</symbol>
                </doStatement>
              </statements>
              <symbol>}</symbol>
            </ifStatement>
            <ifStatement>
              <keyword>if</keyword>
              <symbol>(</symbol>
              <expression>
                <term>
                  <identifier>key</identifier>
                </term>
                <symbol>=</symbol>
                <term>
                  <integerConstant>88</integerConstant>
                </term>
              </expression>
              <symbol>)</symbol>
              <symbol>{</symbol>
              <statements>
                <doStatement>
                  <keyword>do</keyword>
                  <identifier>square</identifier>
                  <symbol>.</symbol>
                  <identifier>incSize</identifier>
                  <symbol>(</symbol>
                  <expressionList></expressionList>
                  <symbol>)</symbol>
                  <symbol>;</symbol>
                </doStatement>
              </statements>
              <symbol>}</symbol>
            </ifStatement>
            <ifStatement>
              <keyword>if</keyword>
              <symbol>(</symbol>
              <expression>
                <term>
                  <identifier>key</identifier>
                </term>
                <symbol>=</symbol>
                <term>
                  <integerConstant>131</integerConstant>
                </term>
              </expression>
              <symbol>)</symbol>
              <symbol>{</symbol>
              <statements>
                <letStatement>
                  <keyword>let</keyword>
                  <identifier>direction</identifier>
                  <symbol>=</symbol>
                  <expression>
                    <term>
                      <integerConstant>1</integerConstant>
                    </term>
                  </expression>
                  <symbol>;</symbol>
                </letStatement>
              </statements>
              <symbol>}</symbol>
            </ifStatement>
            <ifStatement>
              <keyword>if</keyword>
              <symbol>(</symbol>
              <expression>
                <term>
                  <identifier>key</identifier>
                </term>
                <symbol>=</symbol>
                <term>
                  <integerConstant>133</integerConstant>
                </term>
              </expression>
              <symbol>)</symbol>
              <symbol>{</symbol>
              <statements>
                <letStatement>
                  <keyword>let</keyword>
                  <identifier>direction</identifier>
                  <symbol>=</symbol>
                  <expression>
                    <term>
                      <integerConstant>2</integerConstant>
                    </term>
                  </expression>
                  <symbol>;</symbol>
                </letStatement>
              </statements>
              <symbol>}</symbol>
            </ifStatement>
            <ifStatement>
              <keyword>if</keyword>
              <symbol>(</symbol>
              <expression>
                <term>
                  <identifier>key</identifier>
                </term>
                <symbol>=</symbol>
                <term>
                  <integerConstant>130</integerConstant>
                </term>
              </expression>
              <symbol>)</symbol>
              <symbol>{</symbol>
              <statements>
                <letStatement>
                  <keyword>let</keyword>
                  <identifier>direction</identifier>
                  <symbol>=</symbol>
                  <expression>
                    <term>
                      <integerConstant>3</integerConstant>
                    </term>
                  </expression>
                  <symbol>;</symbol>
                </letStatement>
              </statements>
              <symbol>}</symbol>
            </ifStatement>
            <ifStatement>
              <keyword>if</keyword>
              <symbol>(</symbol>
              <expression>
                <term>
                  <identifier>key</identifier>
                </term>
                <symbol>=</symbol>
                <term>
                  <integerConstant>132</integerConstant>
                </term>
              </expression>
              <symbol>)</symbol>
              <symbol>{</symbol>
              <statements>
                <letStatement>
                  <keyword>let</keyword>
                  <identifier>direction</identifier>
                  <symbol>=</symbol>
                  <expression>
                    <term>
                      <integerConstant>4</integerConstant>
                    </term>
                  </expression>
                  <symbol>;</symbol>
                </letStatement>
              </statements>
              <symbol>}</symbol>
            </ifStatement>
            <whileStatement>
              <keyword>while</keyword>
              <symbol>(</symbol>
              <expression>
                <term>
                  <symbol>~</symbol>
                  <term>
                    <symbol>(</symbol>
                    <expression>
                      <term>
                        <identifier>key</identifier>
                      </term>
                      <symbol>=</symbol>
                      <term>
                        <integerConstant>0</integerConstant>
                      </term>
                    </expression>
                    <symbol>)</symbol>
                  </term>
                </term>
              </expression>
              <symbol>)</symbol>
              <symbol>{</symbol>
              <statements>
                <letStatement>
                  <keyword>let</keyword>
                  <identifier>key</identifier>
                  <symbol>=</symbol>
                  <expression>
                    <term>
                      <identifier>Keyboard</identifier>
                      <symbol>.</symbol>
                      <identifier>keyPressed</identifier>
                      <symbol>(</symbol>
                      <expressionList></expressionList>
                      <symbol>)</symbol>
                    </term>
                  </expression>
                  <symbol>;</symbol>
                </letStatement>
                <doStatement>
                  <keyword>do</keyword>
                  <identifier>moveSquare</identifier>
                  <symbol>(</symbol>
                  <expressionList></expressionList>
                  <symbol>)</symbol>
                  <symbol>;</symbol>
                </doStatement>
              </statements>
              <symbol>}</symbol>
            </whileStatement>
          </statements>
          <symbol>}</symbol>
        </whileStatement>
        <returnStatement>
          <keyword>return</keyword>
          <symbol>;</symbol>
        </returnStatement>
      </statements>
      <symbol>}</symbol>
    </subroutineBody>
  </subroutineDec>
  <symbol>}</symbol>
</class>`

			reader := strings.NewReader(inputTokensXml)
			testTokens := tokens.LoadTokensFromXML(reader)

			var buf bytes.Buffer
			xmlEnc := xml.NewEncoder(&buf)

			// Act
			ce, err := New(testTokens, xmlEnc)
			assert.NoError(t, err)
			err = ce.CompileClass()

			// Assert
			assert.NoError(t, err)

			assert.Equal(t, expectedXml, buf.String())
		})
	})
}
