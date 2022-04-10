
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

//go:build (darwin && !cgo) || ios || (linux && arm64) || windows || (!darwin && !freebsd && !linux && !netbsd && !solaris)
// +build darwin,!cgo ios linux,arm64 windows !darwin,!freebsd,!linux,!netbsd,!solaris

// This is the fallback implementation of directory watching.
// It is used on unsupported platforms.

package keystore

type watcher struct{ running bool }

func newWatcher(*accountCache) *watcher { return new(watcher) }
func (*watcher) start()                 {}
func (*watcher) close()                 {}
