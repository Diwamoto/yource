<?php
declare(strict_types=1);

namespace App\Test\TestCase\Model\Table;

use App\Model\Table\UserProfileTable;
use Cake\TestSuite\TestCase;

/**
 * App\Model\Table\UserProfileTable Test Case
 */
class UserProfileTableTest extends TestCase
{
    /**
     * Test subject
     *
     * @var \App\Model\Table\UserProfileTable
     */
    protected $UserProfile;

    /**
     * Fixtures
     *
     * @var array
     */
    protected $fixtures = [
        'app.UserProfile',
        'app.User',
    ];

    /**
     * setUp method
     *
     * @return void
     */
    public function setUp(): void
    {
        parent::setUp();
        $config = $this->getTableLocator()->exists('UserProfile') ? [] : ['className' => UserProfileTable::class];
        $this->UserProfile = $this->getTableLocator()->get('UserProfile', $config);
    }

    /**
     * tearDown method
     *
     * @return void
     */
    public function tearDown(): void
    {
        unset($this->UserProfile);

        parent::tearDown();
    }

    /**
     * Test validationDefault method
     *
     * @return void
     */
    public function testValidationDefault(): void
    {
        $this->markTestIncomplete('Not implemented yet.');
    }

    /**
     * Test buildRules method
     *
     * @return void
     */
    public function testBuildRules(): void
    {
        $this->markTestIncomplete('Not implemented yet.');
    }
}
